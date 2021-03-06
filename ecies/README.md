### 非对称加密算法之 ecies 椭圆曲线数据加解密算法
"crypto/elliptic"
"crypto/ecdsa"

官方实现两个包用于数字签名：
即私钥加密，公钥验签

以太坊基于go官方包"crypto/ecdsa"做了一个ecies算法的实现，用于数据的非对称加解密传输
"github.com/ethereum/go-ethereum/blob/master/crypto/ecies"



### 关于ECC椭圆曲线密码学
- ECDSA是一种数字签名算法
- ECIES是一种集成加密方案
- ECDH是一种关键的安全密钥交换算法

#### 首先，您应该了解这些算法的目的:

##### ECDSA(数字签名算法)用于验证数字内容。
```
有效的数字签名使收件人有理由相信该邮件是由已知发件人创建的，这样发件人就不会拒绝发送邮件（身份验证和不可否认），并且邮件在传输过程中未被更改（完整性）。
```

##### ECIES(集成加密方案)是一种混合加密方案，它为选择的纯文本和选择的密文攻击提供语义安全性。
```
ECIES使用不同类型的功能：
密钥协议功能
关键衍生函数
对称加密方案
哈希函数
```

##### ECDH(安全密钥交换算法)用于通过非安全通道安全地交换密钥。
```
在这里，您对这些算法的Elliptic Curve变体感兴趣。您的要求是交换一些数据。因此，您可以使用ECDH共享密钥，使用ECDSA对内容进行签名。
由于ECDH不提供身份验证，因此我们可以将ECDSA用于此目的。共享密钥后，您可以通过非安全通道安全地交换数据。
可以通过考虑您需要的安全级别和您获得的计算能力来定义密钥的强度。
```