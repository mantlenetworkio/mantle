package abnormal

import (
	"github.com/mantlenetworkio/mantle/mt-tss/node/tsslib/conversion"
)

func (m *Manager) getAbnormalPubKeysInList(peers []string) ([]string, error) {
	var partiesInList []string
	// we convert nodes (in the peers list) P2PID to public key
	for partyID, p2pID := range m.PartyIDtoP2PID {
		for _, el := range peers {
			if el == p2pID.String() {
				partiesInList = append(partiesInList, partyID)
			}
		}
	}

	localPartyInfo := m.partyInfo
	partyIDMap := localPartyInfo.PartyIDMap
	AbnormalPubKeys, err := conversion.AccPubKeysFromPartyIDs(partiesInList, partyIDMap)
	if err != nil {
		return nil, err
	}

	return AbnormalPubKeys, nil
}

func (m *Manager) getAbnormalPubKeysNotInList(peers []string) ([]string, error) {
	var partiesNotInList []string
	// we convert nodes (NOT in the peers list) P2PID to public key
	for partyID, p2pID := range m.PartyIDtoP2PID {
		if m.localPartyID == partyID {
			continue
		}
		found := false
		for _, each := range peers {
			if p2pID.String() == each {
				found = true
				break
			}
		}
		if !found {
			partiesNotInList = append(partiesNotInList, partyID)
		}
	}

	partyIDMap := m.partyInfo.PartyIDMap
	AbnormalPubKeys, err := conversion.AccPubKeysFromPartyIDs(partiesNotInList, partyIDMap)
	if err != nil {
		return nil, err
	}

	return AbnormalPubKeys, nil
}

// GetAbnormalPubKeysNotInList returns the nodes public key who are not in the peer list
func (m *Manager) GetAbnormalPubKeysLists(peer []string) ([]string, []string, error) {
	inList, err := m.getAbnormalPubKeysInList(peer)
	if err != nil {
		return nil, nil, err
	}

	notInlist, err := m.getAbnormalPubKeysNotInList(peer)
	if err != nil {
		return nil, nil, err
	}

	return inList, notInlist, err
}
