import yargs from "yargs";
import { readFileSync } from "fs";
import Orbs from "orbs-client-sdk";
import elliptic from "elliptic";
import sha3 from "js-sha3";

const { keccak256 } = sha3;

function toChecksumAddress(address, chainId = null) {
    if (typeof address !== "string") {
        return "";
    }

    if (!/[0-9a-f]{40}$/i.test(address)) {
        throw new Error(`Given address "${address}" is not a valid Ethereum address.`);
    }

    const stripAddress = address.toLowerCase();
    const prefix = chainId != null ? chainId.toString() + "0x" : "";
    const keccakHash = keccak256(prefix + address)
        .toString("hex")
        .replace(/^0x/i, "")
    let checksumAddress = "";

    for (let i = 0; i < stripAddress.length; i++) {
        checksumAddress += parseInt(keccakHash[i], 16) >= 8 ? stripAddress[i].toUpperCase() : stripAddress[i];
    }

    return checksumAddress;
}

function publicKeyToAddress(publicKey) {
    const rawAddress = Buffer.from(keccak256.arrayBuffer(publicKey).slice(12)).toString("hex");
    return toChecksumAddress(rawAddress);
}

export function main(argv) {
    yargs(argv)
        .command(["node"], "generate a new node address and private key", {}, (argv) => {
            const ec = new elliptic.ec("secp256k1");
            const keyPair = ec.genKeyPair();
            const publicKeyBytes = new Uint8Array(keyPair.getPublic("array")).slice(1);
            const publicKey = Buffer.from(publicKeyBytes).toString("hex");
            const privateKey = keyPair.getPrivate("hex");
            const address = publicKeyToAddress(publicKeyBytes);

            const output = {
                NodePrivateKey: "0x"+privateKey,
                NodePublicKey: "0x"+publicKey,
                NodeAddress: "0x"+address,
            }

            console.log(JSON.stringify(output, 2, 2));
        })
        .command(["client"], "generate a new client address and public/private key pair", {}, async (argv) => {
            const account = await Orbs.createAccount();
            const output = {
                ClientPrivateKey: Orbs.encodeHex(account.privateKey),
                ClientPublicKey: Orbs.encodeHex(account.publicKey),
                ClientAddress: account.address,
            };
            console.log(JSON.stringify(output, 2, 2));
        })
        .command(["version"], "show version", {}, (argv) => {
            const { version } = JSON.parse(readFileSync("./package.json").toString());
            console.log(`orbs-key-generator version ${version}`);
        })
        .demandCommand()
        .help()
        .wrap(80)
        .argv;
}

main(process.argv.slice(2));