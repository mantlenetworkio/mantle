geth --datadir sequencer3 --rpcport 8087 --port 30305 --password password.txt \
--unlock 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC --nodiscover \
--bootnods enode://09998ab210032351a2bea94fbb8c867cfc95553019d3cecfe0d3f3aef78b24290e624b07a99fdaa3c2fb9015d193303179b396e3520284d249ef9a233a5f5e62@172.17.0.1:30303?discport=0 \
--scheduler.address 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266 \
--sequencer.mode=true --verbosity 5 --mine
