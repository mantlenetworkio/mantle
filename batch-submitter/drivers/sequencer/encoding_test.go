package sequencer_test

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/mantlenetworkio/mantle/batch-submitter/drivers/sequencer"
	"github.com/stretchr/testify/require"
)

// TestBatchContextEncodeDecode tests the (de)serialization of a BatchContext
// against the spec test vector. The encoding should be:
//   - num_sequenced_txs:        3 bytes
//   - num_subsequent_queue_txs: 3 bytes
//   - timestamp:                5 bytes
//   - block_number:             5 bytes
func TestBatchContextEncodeDecode(t *testing.T) {
	t.Parallel()

	// Test vector is chosen such that each byte maps one to one with a
	// specific byte of the parsed BatchContext and such that improper
	// choice of endian-ness for any field will fail.
	hexEncoding := "000102030405060708090a0b0c0d0e0f"

	expBatch := sequencer.BatchContext{
		NumSequencedTxs:       0x000102,
		NumSubsequentQueueTxs: 0x030405,
		Timestamp:             0x060708090a,
		BlockNumber:           0x0b0c0d0e0f,
	}

	rawBytes, err := hex.DecodeString(hexEncoding)
	require.Nil(t, err)

	// Test Read produces expected batch.
	var batch sequencer.BatchContext
	err = batch.Read(bytes.NewReader(rawBytes))
	require.Nil(t, err)
	require.Equal(t, expBatch, batch)

	// Test Write produces original test vector.
	var buf bytes.Buffer
	batch.Write(&buf)
	require.Equal(t, hexEncoding, hex.EncodeToString(buf.Bytes()))
}

// AppendSequencerBatchParamsTestCases is an enclosing struct that holds the
// individual AppendSequencerBatchParamsTests. This is the root-level object
// that will be parsed from the JSON, spec test-vectors.
type AppendSequencerBatchParamsTestCases struct {
	Tests []AppendSequencerBatchParamsTest `json:"tests"`
}

// AppendSequencerBatchParamsTest specifies a single instance of a valid
// encode/decode test case for an AppendequencerBatchParams.
type AppendSequencerBatchParamsTest struct {
	Name                  string                   `json:"name"`
	HexEncoding           string                   `json:"hex_encoding"`
	ShouldStartAtElement  uint64                   `json:"should_start_at_element"`
	TotalElementsToAppend uint64                   `json:"total_elements_to_append"`
	Contexts              []sequencer.BatchContext `json:"contexts"`
	Error                 bool                     `json:"error"`
}

var appendSequencerBatchParamTests = AppendSequencerBatchParamsTestCases{}

func init() {
	data, err := os.ReadFile("./testdata/valid_append_sequencer_batch_params.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &appendSequencerBatchParamTests)
	if err != nil {
		panic(err)
	}
}

// TestAppendSequencerBatchParamsEncodeDecode asserts the proper encoding and
// decoding of valid serializations for AppendSequencerBatchParams.
func TestAppendSequencerBatchParamsEncodeDecode(t *testing.T) {
	t.Parallel()

	for _, test := range appendSequencerBatchParamTests.Tests {
		t.Run(test.Name, func(t *testing.T) {
			testAppendSequencerBatchParamsEncodeDecode(t, test)
		})
	}
}

func testAppendSequencerBatchParamsEncodeDecode(
	t *testing.T, test AppendSequencerBatchParamsTest) {
	// Construct the params we expect to decode, minus the txs. Those are
	// compared separately below.
	expParams := sequencer.AppendSequencerBatchParams{
		ShouldStartAtElement:  test.ShouldStartAtElement,
		TotalElementsToAppend: test.TotalElementsToAppend,
		Contexts:              test.Contexts,
	}

	// Decode the batch from the test string.
	rawBytes, err := hex.DecodeString(test.HexEncoding)
	require.Nil(t, err)

	var params sequencer.AppendSequencerBatchParams
	err = params.Read(bytes.NewReader(rawBytes))
	if test.Error {
		require.ErrorIs(t, err, sequencer.ErrMalformedBatch)
	} else {
		require.Nil(t, err)
	}

	// Assert that the decoded params match the expected params. The
	// transactions are compared serparetly (via hash), since the internal
	// `time` field of each transaction will differ. This field is only used
	// for spam prevention, so it is safe to ignore wrt. to serialization.
	// The decoded txs are reset on the the decoded params afterwards to
	// test the serialization.
	require.Equal(t, expParams, params)

	// Finally, encode the decoded object and assert it matches the original
	// hex string.
	paramsBytes, err := params.Serialize(sequencer.BatchTypeZlib)

	// Return early when testing error cases, no need to reserialize again
	if test.Error {
		require.ErrorIs(t, err, sequencer.ErrMalformedBatch)
		return
	}

	require.Nil(t, err)
	require.Equal(t, test.HexEncoding, hex.EncodeToString(paramsBytes))

	// Serialize the batches in compressed form
	compressedParamsBytes, err := params.Serialize(sequencer.BatchTypeZlib)
	require.Nil(t, err)

	// Deserialize the compressed batch
	var paramsCompressed sequencer.AppendSequencerBatchParams
	err = paramsCompressed.Read(bytes.NewReader(compressedParamsBytes))
	require.Nil(t, err)

	require.Equal(t, expParams, paramsCompressed)
}

// TestMarkerContext asserts that each batch type returns the correct marker
// context.
func TestMarkerContext(t *testing.T) {
	batchTypes := []sequencer.BatchType{
		sequencer.BatchTypeLegacy,
		sequencer.BatchTypeZlib,
	}

	for _, batchType := range batchTypes {
		t.Run(batchType.String(), func(t *testing.T) {
			markerContext := batchType.MarkerContext()
			if batchType == sequencer.BatchTypeLegacy {
				require.Nil(t, markerContext)
			} else {
				require.NotNil(t, markerContext)

				// All marker contexts MUST have a zero timestamp.
				require.Equal(t, uint64(0), markerContext.Timestamp)

				// Currently all other fields besides block number are defined
				// as zero.
				require.Equal(t, uint64(0), markerContext.NumSequencedTxs)
				require.Equal(t, uint64(0), markerContext.NumSubsequentQueueTxs)

				// Assert that the block number for each batch type is set to
				// the correct constant.
				switch batchType {
				case sequencer.BatchTypeZlib:
					require.Equal(t, uint64(0), markerContext.BlockNumber)
				default:
					t.Fatalf("unknown batch type")
				}

				// Ensure MarkerBatchType produces the expected BatchType.
				require.Equal(t, batchType, markerContext.MarkerBatchType())
			}
		})
	}
}

// TestIsMarkerContext asserts that IsMarkerContext returns true iff the
// timestamp is zero.
func TestIsMarkerContext(t *testing.T) {
	batchContext := sequencer.BatchContext{
		NumSequencedTxs:       1,
		NumSubsequentQueueTxs: 2,
		Timestamp:             3,
		BlockNumber:           4,
	}
	require.False(t, batchContext.IsMarkerContext())

	batchContext = sequencer.BatchContext{
		NumSequencedTxs:       0,
		NumSubsequentQueueTxs: 0,
		Timestamp:             3,
		BlockNumber:           0,
	}
	require.False(t, batchContext.IsMarkerContext())

	batchContext = sequencer.BatchContext{
		NumSequencedTxs:       1,
		NumSubsequentQueueTxs: 2,
		Timestamp:             0,
		BlockNumber:           4,
	}
	require.True(t, batchContext.IsMarkerContext())
}

func TestReadUint64(t *testing.T) {
	readUint64 := func(r io.Reader, val *uint64, n uint) error {
		var byteOrder = binary.BigEndian
		var buf [8]byte
		if n > 8 {
			return fmt.Errorf("bytes shift out of range")
		}
		if _, err := r.Read(buf[8-n:]); err != nil {
			return err
		}
		*val = byteOrder.Uint64(buf[:])
		return nil
	}

	var be = make([]byte, 8)
	var x uint64
	bytes.NewBuffer(be)
	binary.BigEndian.PutUint64(be, 100)
	require.NoError(t, readUint64(bytes.NewBuffer(be), &x, 8))
	require.Equal(t, x, uint64(100))

	require.Error(t, readUint64(bytes.NewBuffer(be), &x, 9))
}
