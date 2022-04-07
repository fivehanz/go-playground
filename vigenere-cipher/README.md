# vigenere cipher

implemented with utf8 in mind, can encrypt and decrypt utf8 characters (moslty.)

## build

```
╭─ ~/Projects/go-playground/vigenere-cipher on main +2 ·································································································································································································· Py base at 23:55:00
╰─❯ go build                                                                                                                                                                                                                                              

╭─ ~/Projects/go-playground/vigenere-cipher on main +2 ?1 ······························································································································································································· Py base at 23:55:03
╰─❯ ls
.rw-r--r--   66 hz 20 Mar 04:15 go.mod
.rw-r--r-- 3.6k hz  7 Apr 23:50 main.go
.rw-r--r-- 5.5k hz  7 Apr 23:54 README.md
.rwxr-xr-x 2.0M hz  7 Apr 23:55 vigenere-cipher

╭─ ~/Projects/go-playground/vigenere-cipher on main +2 !1 ?1 ···························································································································································································· Py base at 23:55:58
╰─❯ ./vigenere-cipher 
usage:
$ vigenere-cipher KEY MESSAGE
sample test case below:

Sample-Key: Lorem Isum ßåßˆøœ∆˚å∫ß∂˙∫∫ßåç≈
Sample-Text: Eg: Go says so, that! A@B#C$- aå 水目馬 ;). πø∑ß∂œ∫√∆¥

Encrypted Text: Ö¬´iæÖæÿŘŎ˲ĘǇ≮̻ř≌ÿ≃̙≭≎ĢĉĔ≨­Ŕ沙睛駌i®@ҟǝ⋰Υ⋺ʦ䐱⓴⋫⋐
Decrypted Text: Eg: Go says so, that! A@B#C$- aå 水目馬 ;). πø∑ß∂œ∫√∆¥
```

## examples

### default run

```
$ go run main.go            
usage:
$ vigenere-cipher KEY MESSAGE
sample test case below:

Sample-Key: Lorem Isum ßåßˆøœ∆˚å∫ß∂˙∫∫ßåç≈
Sample-Text: Eg: Go says so, that! A@B#C$- aå 水目馬 ;). πø∑ß∂œ∫√∆¥

Encrypted Text: Ö¬´iæÖæÿŘŎ˲ĘǇ≮̻ř≌ÿ≃̙≭≎ĢĉĔ≨­Ŕ沙睛駌i®@ҟǝ⋰Υ⋺ʦ䐱⓴⋫⋐
Decrypted Text: Eg: Go says so, that! A@B#C$- aå 水目馬 ;). πø∑ß∂œ∫√∆¥

```
### provide only key

```
$ go run main.go XYZ
Key: XYZ
Sample-Text: Eg: Go says so, that! A@B#C$- aå 水目馬 ;). πø∑ß∂œ∫√∆¥

Encrypted Text: Àx ÉxÌ»ÑÌzËÈxÍÂ¹Í{x||z¹ľz沌睇騆xyКŐ≪Ĺ≚Ƭ⊅≲≟ÿ
Decrypted Text: Eg: Go says so, that! A@B#C$- aå 水目馬 ;). πø∑ß∂œ∫√∆¥
```
### provide key and message

```
╭─ ~/Projects/go-playground/vigenere-cipher on main +1 !2 ······························································································································································································· Py base at 23:50:06
╰─❯ go run main.go key message
Key: key
Text: message

Encrypted Text: ØÊìÞÆàÐ
Decrypted Text: message

╭─ ~/Projects/go-playground/vigenere-cipher on main +1 !2 ······························································································································································································· Py base at 23:50:10
╰─❯ go run main.go key MESSAGE
Key: key
Text: MESSAGE

Encrypted Text: ¸ªÌ¾¦À°
Decrypted Text: MESSAGE

╭─ ~/Projects/go-playground/vigenere-cipher on main +1 !2 ······························································································································································································· Py base at 23:53:12
╰─❯ go run main.go "sample big key 1" "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin pretium lacus in tellus placerat ultricies. Praesent accumsan rhoncus justo in posuere. Fusce odio elit, hendrerit sed imperdiet a, fringilla quis."
Key: sample big key 1
Text: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin pretium lacus in tellus placerat ultricies. Praesent accumsan rhoncus justo in posuere. Fusce odio elit, hendrerit sed imperdiet a, fringilla quis.

Encrypted Text: ¿ÐßÕÙÒÜÜÉè åàÙàÏÎÛLÈè¤ØÄáÕàÚÊËÛÎìáÈÕØÎ·ÚÎç@¡åÆáÙáÒ@ÎÊÊÞâQçÆÙÜáØ@ÒÕÈÐ×ÚQèÍáâÕÈÇÜ@»×Ú¤ØÏáÍÈ×ÖÚÙë áÄâãÏÕÝÖ@ÔÓ æÖÒâÑ@¨ÞÚÐèâÒÜÕÙLÑÌÏ×ÞçàÕÐÏÙÌÏÎÞQÔÖÞÎÉÒÓÌêæ
Decrypted Text: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin pretium lacus in tellus placerat ultricies. Praesent accumsan rhoncus justo in posuere. Fusce odio elit, hendrerit sed imperdiet a, fringilla quis.

╭─ ~/Projects/go-playground/vigenere-cipher on main +1 !2 ······························································································································································································· Py base at 23:53:20
╰─❯ go run main.go "sample BIG key 2" "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin pretium lacus in tellus placerat ultricies. Praesent accumsan rhoncus justo in posuere. Fusce odio elit, hendrerit sed imperdiet a, fringilla quis."
Key: sample BIG key 2
Text: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin pretium lacus in tellus placerat ultricies. Praesent accumsan rhoncus justo in posuere. Fusce odio elit, hendrerit sed imperdiet a, fringilla quis.

Encrypted Text: ¿ÐßÕÙ²¼¼Éè¡åàÙà¯®»LÈè¥ØÄáÕàÚbª«ÛÎìáÈÕØÎpiÚÎç@¢åÆáÙáÒ@®ªªÞâRçÆÙÜáØ@²µ¨Ð×ÚRèÍáâÕÈ§¼u@»×Ú¥ØÏáÍÈ·¶ºÙë¡áÄâãÏµ½¶@ÔÓ¡æÖÒâÑ@¾ºÐèâÒÜÕÙLb±¬Ï×ÞçàÕÐ¯¹¬ÏÎÞRÔÖÞÎ©²³Ìêæ
Decrypted Text: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin pretium lacus in tellus placerat ultricies. Praesent accumsan rhoncus justo in posuere. Fusce odio elit, hendrerit sed imperdiet a, fringilla quis.

```

# To Do 

- [x] abort ASCII 
- [x] try with "unicode/uft8" package
- [x] uses rune; golang is by default in utf8 
