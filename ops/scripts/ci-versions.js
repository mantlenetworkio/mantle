const os = require('os')

// this script unbundles the published packages output
// from changesets action to a key-value pair to be used
// with our publishing CI workflow
data = process.argv[2]
data = JSON.parse(data)

for (const i of data) {
<<<<<<< HEAD
  const name = i.name.replace("@mantlenetworkio/", "")
=======
  const name = i.name.replace("@eth-optimism/", "")
>>>>>>> op-node/v1.1.1
  const version = i.version
  process.stdout.write(`::set-output name=${name}::${version}` + os.EOL)
}
