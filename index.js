import yargs from "yargs";
import { readFileSync } from "fs";
import { generateNodeKeys, generateOrbsKeys } from "./keys.js";

export function main(argv) {
    yargs(argv)
        .command(["node"], "generate a new node address and private key", {}, (argv) => {
            const output = generateNodeKeys();
            console.log(JSON.stringify(output, 2, 2));
        })
        .command(["client"], "generate a new client address and public/private key pair", {}, async (argv) => {
            const output = await generateOrbsKeys();
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