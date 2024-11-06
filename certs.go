package main

var TLSCertKey = []byte(`-----BEGIN CERTIFICATE-----
MIIF5TCCBM2gAwIBAgIQfbCIu2TrNWkekaTIv0FesjANBgkqhkiG9w0BAQsFADAh
MQ4wDAYDVQQDDAVteXNxbDEPMA0GA1UECgwGa3ViZWRiMB4XDTI0MDkxMzExNDIx
NloXDTI0MTIxMjExNDIxNlowKTEnMCUGA1UEAxMeZHJ1aWQtcXVpY2tzdGFydC1w
b2RzLmRlbW8uc3ZjMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAueJN
EVuGWGYiauPNANS0PgDxDAd3od4+dh8XI0ZJ8phImBYCIehUhvtJbpkv4Vw0hkuU
XauL/l4A7s5Lvb4zqSQDkFWVCi2ui1dBc3xvJF/dxKrRT1RCT/3LeFx0eK2YGJC9
MlC3TKJ/ZMJ+9ctJHFoXn9nR+k2n6VQQKTToEvSEPoQd+BEBIEbBSuOoE6vucSNi
yM06wlQCy8610gOy4ThZhKW0Y9pN2distEB/u4lF7/Go4bLqviJ6HAFG2U4ZyVZ+
xUp8U0kD3Jh37BSoeOIhC9Cfh+Z5/3/YrkUd8oIphxlOtdDviYYRvaWCeF7YNoxo
kUns3xycyu8zn0jasQIDAQABo4IDDzCCAwswDgYDVR0PAQH/BAQDAgKkMBMGA1Ud
JQQMMAoGCCsGAQUFBwMCMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFHbckBFc
aLnhyERbFOqvsLHoin7BMB8GA1UdIwQYMBaAFEjs4Zz1/lrmAICyffSwEdQpvlBu
MIICkQYDVR0RBIICiDCCAoSCLiouZHJ1aWQtcXVpY2tzdGFydC1wb2RzLmRlbW8u
c3ZjLmNsdXN0ZXIubG9jYWyCTGRydWlkLXF1aWNrc3RhcnQtYnJva2Vycy0wLmRy
dWlkLXF1aWNrc3RhcnQtcG9kcy5kZW1vLnN2Yy5jbHVzdGVyLmxvY2FsOjgwODKC
UWRydWlkLXF1aWNrc3RhcnQtY29vcmRpbmF0b3JzLTAuZHJ1aWQtcXVpY2tzdGFy
dC1wb2RzLmRlbW8uc3ZjLmNsdXN0ZXIubG9jYWw6ODA4MYJQZHJ1aWQtcXVpY2tz
dGFydC1oaXN0b3JpY2Fscy0wLmRydWlkLXF1aWNrc3RhcnQtcG9kcy5kZW1vLnN2
Yy5jbHVzdGVyLmxvY2FsOjgwODOCU2RydWlkLXF1aWNrc3RhcnQtbWlkZGxlbWFu
YWdlcnMtMC5kcnVpZC1xdWlja3N0YXJ0LXBvZHMuZGVtby5zdmMuY2x1c3Rlci5s
b2NhbDo4MDkxghVkcnVpZC1xdWlja3N0YXJ0LXBvZHOCHmRydWlkLXF1aWNrc3Rh
cnQtcG9kcy5kZW1vLnN2Y4IsZHJ1aWQtcXVpY2tzdGFydC1wb2RzLmRlbW8uc3Zj
LmNsdXN0ZXIubG9jYWyCTGRydWlkLXF1aWNrc3RhcnQtcm91dGVycy0wLmRydWlk
LXF1aWNrc3RhcnQtcG9kcy5kZW1vLnN2Yy5jbHVzdGVyLmxvY2FsOjg4ODiCCWxv
Y2FsaG9zdIIRbXktZ3JvdXAuZGVtby5zdmOCIW15bWluaW8taGwuZGVtby5zdmMu
Y2x1c3Rlci5sb2NhbIIQemstZGVtby5kZW1vLnN2Y4cEfwAAATANBgkqhkiG9w0B
AQsFAAOCAQEAWKpAjKAr0+gIMCMxPvIEYdyijTLkmsBtRZewq0kKWq5Qc3OgCQ2u
+vPgNYO0zm9x6jKZJTNy988zHZlrsBUB+LOkAAlRK/gLlf/kZiHwSmkraNlxcd1J
1LkznyyVsK/jXirW5N2Hnw0CqEGJHHjslbKt2IxNXEMAeJ8Q73u2qJXPeLo3C2O1
R33qGwXyXowin2GZHefTB2w3vtv9lENXelU4l1XV2vwEnairJQM+rUiLA4216FuN
1lurh1X5cKYCp3yQR/4U9gCbUuGjZqFUyZpWLKQbiN60p2CG6g7Iy441sRRtFaW3
zEiQ/mOrp75IDfSbmxXEeFBbuoXC8VHCRA==
-----END CERTIFICATE-----`)

var TLSPrivateKeyKey = []byte(`-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC54k0RW4ZYZiJq
480A1LQ+APEMB3eh3j52HxcjRknymEiYFgIh6FSG+0lumS/hXDSGS5Rdq4v+XgDu
zku9vjOpJAOQVZUKLa6LV0FzfG8kX93EqtFPVEJP/ct4XHR4rZgYkL0yULdMon9k
wn71y0kcWhef2dH6TafpVBApNOgS9IQ+hB34EQEgRsFK46gTq+5xI2LIzTrCVALL
zrXSA7LhOFmEpbRj2k3Z2Ky0QH+7iUXv8ajhsuq+InocAUbZThnJVn7FSnxTSQPc
mHfsFKh44iEL0J+H5nn/f9iuRR3ygimHGU610O+JhhG9pYJ4Xtg2jGiRSezfHJzK
7zOfSNqxAgMBAAECggEAeGs6Mfxw0+OWeFiK/fbNrT0+6uQLhDlNjYgmnRtz3Ez6
+kuHph+0Yp13vMZQVlY9UKK1wtjXIDoVsJ7EQm9tWC2UgXRv0OA3Kt2j8QSzm4Tf
TgC7CAE11Ew9/AMpEOSkm4Voi6fulx7OGKBHT0QaYNucYlnhxZ15IZbrwUYc8emK
9IoR0ieNjE69N/8wSPZEuCC6SnPOWElN4SLObBnt/r5XGluFhYXq1iipqZMIQM46
qq8ZPcdfTNtTdqvNZFavMmBWyS7RNDuWjvi+p03e8IQ5VuvOkjb3L2TpVcVH85MA
ephiP/Dn6Yw+Zz7UouNxuC/KUyDggf/nGNmXF4NegQKBgQDWc5bxye7j3DCA7q4B
DD6scpaPQDbc/HIkv4qKEUJsIGW4yEhixvZDnTSqRE2+MrXnqTdKkA5Oqttm4rVG
UKGVMu6dkXVMEXJz/bNxuT/pDqYywq+7nCCS89X8Ek8Sf2DXjiYiR0eaHwQTLZ9L
SleAmtMLfCbyJvLmYQ4T6YvJjQKBgQDd5dAkoSxLW776tj5hwhGq37Fsq6sB5x7Z
+y+TbGJhFlKqPAj/y2t9mReW2kj+rWY7ig7PT8INdS9vu4tc0Aq0rLHSYllMaO4u
jxCMo1fq+9iY4pdg3qrIwwpVbf+cEAf4gVBjhI/xc6XiRiPbzmuNqVFxhJiDjcJU
3xgRR7FCtQKBgQCIWnqRRjoUu35DnH8av3RiJYHBl6zKmeDhqMAnxY2cPoUvFnmg
BK20lBJWaJOd6ZRrtdYoHKxcLJrQMjnceYwj13yMx13zfext/9PG7WjLgLr/73XZ
Lg4wIcXfPqz0L/WdWQ66IORQBISMxSdRsfUtGkQyO0BwG3+6J9/RnV96SQKBgHok
5aU0HvjAYOkK5l3TgPwpNDYYbQKYIXBmBzNEZwcATbKtZ1q+s8WoPsboE2Q1OOAv
R6WwRqY3ykvb1YPadPotUZj4UCVyYSG52KdBxsRvqzPFDwTgrOvkCM0rxpc2b9zH
+eM99io3ualLaLO29Zc+C7nL/lPz9XPYdhwU9as9AoGATA/Jkcl+WKV3ZME5R7jM
notEafT9KWQc2cgiZWeF4XSNN0bVTpwE9McBEGBXNpnF62vvMri+kXNuuxOVKhgL
vZdFunuzLuAZUhVDp8FBl7oiDT0salBz1CjYWGZVXWs8fnHFoPqtQx1Y8Jrg8RFp
eTslFjy2vTjITkGPXi+kG1k=
-----END PRIVATE KEY-----`)

var TLSCACertFileName = []byte(`-----BEGIN CERTIFICATE-----
MIIDIzCCAgugAwIBAgIUREO/E8TPnX8oqpDquX0aiY/PVacwDQYJKoZIhvcNAQEL
BQAwITEOMAwGA1UEAwwFbXlzcWwxDzANBgNVBAoMBmt1YmVkYjAeFw0yNDA5MTMx
MDM5MDdaFw0yNTA5MTMxMDM5MDdaMCExDjAMBgNVBAMMBW15c3FsMQ8wDQYDVQQK
DAZrdWJlZGIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC4HvfyJU93
/6XS5brWqXQif6nNTKHuBkkCs0gtfn2lQNrgbX0g9aM7qP+N/8AcFU08JOlJmnRN
ggCtnR95lLvVGYRBWHpBOBEIlGddxu6aqV3MHMuWSUutK57ofxVOe+7iaMnuKDfd
AHTPx8libtf7MuvtI5JEC+vN+GoEI3tE3m+W8ppFVNs4Wm75P9WuUikrgolqwruW
klG8WhmYg4J7gEQdcOG2EJoV3BZmhk1ViliOGHbi04AWgT5JxoQzMtqfJUSMPeh0
EWTTLL9S5DBgygm/bOijQAAJSo+Y9GyON7x027O1cq0wdogIGqVq2yWto2irbMDD
iVJtpX4W50pLAgMBAAGjUzBRMB0GA1UdDgQWBBRI7OGc9f5a5gCAsn30sBHUKb5Q
bjAfBgNVHSMEGDAWgBRI7OGc9f5a5gCAsn30sBHUKb5QbjAPBgNVHRMBAf8EBTAD
AQH/MA0GCSqGSIb3DQEBCwUAA4IBAQAj68LsRrmdXWMh6X8iamz8+c1736t5VHRW
6bnqrarIw2zOr7OpU9omlZlM83LcZcTsuiFVLWMmpknxC5CB/zQsaEOb7nf0Q39/
ZY5PKuEeYPSKc/CoXmrNauiQU0kqnt0qTFKz6YLyTyrvVYsoc95ISp66qFLLIZk9
BpeMTYJa33nv5kxilsYUGn/BZBsahY4Mzh1zAre4jaDxpBvvF/+Y/0RdQZQm4JWQ
FIpDCSa+0+gNP96I0vSTtymIFL+CKdZVT3oL7ErTkzo65a2E9sR1dK66wVJse3+j
uKwu3Edl3q8D/HfL6n+aa+aWf4M8VYDD9KaophWCrPbrIoD41Sdl
-----END CERTIFICATE-----`)