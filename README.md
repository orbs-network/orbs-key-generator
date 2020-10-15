# Orbs Node and Client Key/Address Generator
This toolkit provides you with a simple way to generate both client key pairs and addresses and node key pairs and addresses.

Since the keys are sensitive, it is recommended to use these tools in a private place, where the generated output can be safely stored in a secure location.

## Installing

To install the key generator, use npm

    npm i -g orbs-key-generator

It is also possible to clone this repository and then run `build.sh`

## Generating a client key

To generate a client public and private key, as well as an Orbs address, run:

    orbs-key-generator client
    
## Generating a node key

To generate a node public and private key, as well as an Orbs node address, run:

    orbs-key-generator node
    
## Testing

```
npm i
npm test
```
