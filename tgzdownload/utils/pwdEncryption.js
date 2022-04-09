const CryptoJS = require('crypto-js')


// 长度128bit,192bit,256bit之一
const key = CryptoJS.enc.Utf8.parse('1234567890000000');
const iv = CryptoJS.enc.Utf8.parse("1234567890000000");

// 加密
const encrypt = (word) => {
    const srcs = CryptoJS.enc.Utf8.parse(word);
    const encrypted = CryptoJS.AES.encrypt(srcs, key, {
        iv: iv,
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7
    });
    return encrypted.ciphertext.toString();
}

// 解密
const decrypt = (word) => {
    const encryptedHexStr = CryptoJS.enc.Hex.parse(word);
    const srcs = CryptoJS.enc.Base64.stringify(encryptedHexStr);
    const decrypt = CryptoJS.AES.decrypt(srcs, key, {
        iv: iv,
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7
    });
    const decryptedStr = decrypt.toString(CryptoJS.enc.Utf8);
    return decryptedStr.toString();
}

export default {
    encrypt,
    decrypt,
}