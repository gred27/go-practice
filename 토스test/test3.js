function parseSearch(search) {
    /* 쿼리 문자열 `search`를 파싱하는 함수를 작성해주세요. */
    if (!search) {
        return {};
    }
    const parseStr = search.replace("?", "").split("&");
    const dic = {};

    parseStr.forEach((str) => {
        const [key, value] = str.split("=");
        if (!dic[key]) {
            dic[key] = [value];
        } else {
            dic[key].push(value);
        }
    });

    for (const k in dic) {
        if (dic[k].length === 1) {
            dic[k] = dic[k][0];
        }
    }

    return dic;
}

/*
 * NOTE: 아래 코드는 코드 동작을 확인하기 위한 코드입니다. 수정하지 말아주세요.
 */
function solution(search) {
    var query = parseSearch(search);
    return submit(query);
}

function submit(answer) {
    return JSON.stringify(answer);
}
