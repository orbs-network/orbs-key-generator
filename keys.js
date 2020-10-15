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

export function publicKeyToAddress(publicKey) {
    const rawAddress = Buffer.from(keccak256.arrayBuffer(publicKey).slice(12)).toString("hex");
    return toChecksumAddress(rawAddress);
}

export function generateNodeKeys() {
    const ec = new elliptic.ec("secp256k1");
    const keyPair = ec.genKeyPair();
    const publicKeyBytes = new Uint8Array(keyPair.getPublic("array")).slice(1);
    const publicKey = Buffer.from(publicKeyBytes).toString("hex");
    const privateKey = keyPair.getPrivate("hex");
    const address = publicKeyToAddress(publicKeyBytes);

    return {
        NodePrivateKey: "0x" + privateKey,
        NodePublicKey: "0x" + publicKey,
        NodeAddress: "0x" + address,
    }
}

export async function generateOrbsKeys() {
    const account = await Orbs.createAccount();
    return {
        ClientPrivateKey: Orbs.encodeHex(account.privateKey),
        ClientPublicKey: Orbs.encodeHex(account.publicKey),
        ClientAddress: account.address,
    };
}