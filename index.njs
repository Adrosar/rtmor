#!/usr/bin/env node

"use strict";

var path = require('path');
var cp = require('child_process');

var DICT = {
    'win32-x64': 'build/windows-amd64/rtmor.exe',
    'win32-x32': 'build/windows-386/rtmor.exe',
    'darwin-x64': 'build/darwin-amd64/rtmor',
    'linux-x64': 'build/linux-amd64/rtmor',
    'linux-x32': 'build/linux-386/rtmor'
}

function getArgv() {
    for (var i = 0; i < process.argv.length; i++) {
        var arg = process.argv[i];
        if (arg.indexOf('index.njs') > -1) {
            return process.argv.slice(i + 1);
        }
    }

    return process.argv;
}

function main() {
    var key = process.platform + "-" + process.arch;
    var value = DICT[key];
    if (value) {
        var exec = path.resolve(__dirname, value)
        cp.spawn(exec, getArgv(), {
            cwd: path.resolve(__dirname),
            detached: false,
            stdio: [0, 1, 2]
        });
    } else {
        console.error("Your system \"" + process.platform + "-" + process.arch + "\" is NOT supported!");
    }
}

if (typeof process === "object") {
    main();
}