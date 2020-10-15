import yargs from "yargs";
import { readFileSync } from "fs";
import Orbs from "orbs-client-sdk";

export function main(argv) {
    yargs(argv)
        .command(["node"], "generate a new node address and private key", {}, (argv) => {
            console.log("not supported");
        })
        .command(["client"], "generate a new client address and public/private key pair", {}, async (argv) => {
            const account = await Orbs.createAccount();
            const output = {
                ClientPrivateKey: Orbs.encodeHex(account.privateKey),
                ClientPublicKey: Orbs.encodeHex(account.publicKey),
                ClientAddress: account.address,
            };
            console.log(output);
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