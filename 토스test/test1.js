/*
 * `codeOwnersMap`과 `directory`를 입력받아
 * `directory`의 코드 주인 목록을 반환하는 함수를 작성하세요.
 */
function solution(codeOwnersMap, directory) {
    const directoryList = directory.split("/");
    let curDir = codeOwnersMap;

    if (directoryList.length === 1) {
        return codeOwnersMap[directoryList[0]];
    }

    for (let i = 0; i < directoryList.length; i++) {
        if (!curDir[directoryList[i]]) {
            return curDir[directoryList[i - 1]];
        }

        curDir = curDir[directoryList[i]];
    }

    return curDir;
}

console.log(
    solution(
        {
            scripts: ["배수진"],
            services: {
                "business-ledger": ["고찬균", "배수진"],
                "toss-card": ["채주민", "유재섭"],
                subsidy: ["허주엽"],
                payments: ["유재섭"],
            },
        },
        // "scripts"
        "services/toss-card"
    )
);

//   libraries: {
//     'yarn-workspaces-plugin-since': [ '유재섭', '배수진' ],
//     tds: [ '유병솔', '강두한' ]
//   }
