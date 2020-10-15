import { generateNodeKeys, generateOrbsKeys, publicKeyToAddress } from "../keys.js";
import chai from "chai";
import elliptic from "elliptic";
const { expect } = chai;

function _verifyKeys(address, privateKey) {
    const ec = new elliptic.ec('secp256k1');
    const keyPair = ec.keyFromPrivate(privateKey);
    const publicKey = new Uint8Array(keyPair.getPublic("array")).slice(1);

    const nodeAddress = publicKeyToAddress(publicKey);
    return nodeAddress.toLowerCase() == address.toLowerCase();
}

describe("Keys", () => {
    describe("#generateOrbsKeys", () => {
        it("should conform to a format", async () => {
            const output = await generateOrbsKeys();

            expect(output.ClientAddress).not.to.be.empty;
            expect(output.ClientPrivateKey).not.to.be.empty;
            expect(output.ClientPublicKey).not.to.be.empty;
        });
    });

    describe("#generateNodeKeys", () => {
        it("should conform to a format", async () => {
            const output = generateNodeKeys();

            expect(output.NodeAddress).not.to.be.empty;
            expect(output.NodePrivateKey).not.to.be.empty;
            expect(output.NodePublicKey).not.to.be.empty;
        });

        it("should produce valid keys", () => {
            const output = generateNodeKeys();

            expect(_verifyKeys(output.NodeAddress, output.NodePrivateKey));
        })
    });
});