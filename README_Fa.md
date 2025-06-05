# XrayR

[![](https://img.shields.io/badge/TgChat-@XrayR讨论-blue.svg)](https://t.me/XrayR_project)
[![](https://img.shields.io/badge/Channel-@XrayR通知-blue.svg)](https://t.me/XrayR_channel)
![](https://img.shields.io/github/stars/Aqr-K/XrayR)
![](https://img.shields.io/github/forks/Aqr-K/XrayR)
![](https://github.com/Aqr-K/XrayR/actions/workflows/release.yml/badge.svg)
![](https://github.com/Aqr-K/XrayR/actions/workflows/docker.yml/badge.svg)
[![Github All Releases](https://img.shields.io/github/downloads/Aqr-K/XrayR/total.svg)]()

[Iranian(farsi) README](https://github.com/Aqr-K/XrayR/blob/master/README_Fa.md), [Vietnamese(vi) README](https://github.com/Aqr-K/XrayR/blob/master/README-vi.md), [English(en) README](https://github.com/Aqr-K/XrayR/blob/master/README-en.md)

یک فریمورک بک اند مبتنی بر xray که از چند از پنل پشتیبانی می کند

یک چارچوب بک‌اند مبتنی بر Xray که از پروتکل‌های V2ay، Trojan و Shadowsocks پشتیبانی می‌کند، به راحتی قابل گسترش است و از اتصال چند پنل پشتیبانی می‌کند.

اگر این پروژه را دوست دارید، می توانید با کلیک بر روی ستاره+ساعت در گوشه بالا سمت راست به ادامه روند پیشرفت این پروژه توجه کنید.

آموزش：[اموزش با جزئیات](https://Aqr-K.github.io/XrayR-doc/)

## سلب مسئولیت

این پروژه فقط مطالعه، توسعه و نگهداری شخصی من است. من هیچ گونه قابلیت استفاده را تضمین نمی کنم و مسئولیتی در قبال عواقب ناشی از استفاده از این نرم افزار ندارم.

## امکانات

- منبع باز دائمی و رایگان
- پشتیبانی از چندین پروتکل V2ray، Trojan، Shadowsocks.
- پشتیبانی از ویژگی های جدید مانند Vless و XTLS.
- پشتیبانی از اتصال یک نمونه چند پانل، چند گره، بدون نیاز به شروع مکرر.
- پشتیبانی محدود IP آنلاین
- پشتیبانی از سطح پورت گره، محدودیت سرعت سطح کاربر.
- پیکربندی ساده و سرراست است.
- پیکربندی را تغییر دهید تا نمونه به طور خودکار راه اندازی مجدد شود.
- کامپایل و ارتقاء آن آسان است و می تواند به سرعت نسخه اصلی را به روز کند و از ویژگی های جدید Xray-core پشتیبانی می کند.

## امکانات

| امکانات                                   | v2ray | trojan | shadowsocks |
| ----------------------------------------- | ----- | ------ | ----------- |
| اطلاعات گره را دریافت کنید                | √     | √      | √           |
| دریافت اطلاعات کاربر                      | √     | √      | √           |
| آمار ترافیک کاربران                       | √     | √      | √           |
| گزارش اطلاعات سرور                        | √     | √      | √           |
| به طور خودکار برای گواهی tls درخواست دهید | √     | √      | √           |
| تمدید خودکار گواهی tls                    | √     | √      | √           |
| آمار آنلاین                               | √     | √      | √           |
| محدودیت کاربر آنلاین                      | √     | √      | √           |
| قوانین حسابرسی                            | √     | √      | √           |
| محدودیت سرعت پورت گره                     | √     | √      | √           |
| محدودیت سرعت بر اساس کاربر                | √     | √      | √           |
| DNS سفارشی                                | √     | √      | √           |

## پشتیبانی از قسمت فرانت

| قسمت فرانت                                                  | v2ray | trojan | shadowsocks                           |
| ----------------------------------------------------------- | ----- | ------ | ------------------------------------- |
| sspanel-uim                                                 | √     | √      | √ (تک پورت چند کاربره و V2ray-Plugin) |
| v2board                                                     | √     | √      | √                                     |
| [PMPanel](https://github.com/ByteInternetHK/PMPanel)        | √     | √      | √                                     |
| [ProxyPanel](https://github.com/ProxyPanel/ProxyPanel)      | √     | √      | √                                     |
| [WHMCS (V2RaySocks)](https://v2raysocks.doxtex.com/)        | √     | √      | √                                     |
| [BunPanel](https://github.com/pennyMorant/bunpanel-release) | √     | √      | √                                     |

## نصب نرم افزار

### نصب بصورت یکپارچه

```
wget -N https://raw.githubusercontent.com/Aqr-K/XrayR-release/master/install.sh && bash install.sh
```

### استقرار نرم افزار با استفاده از Docker

[آموزش استقرار داکر](https://Aqr-K.github.io/XrayR-doc/xrayr-xia-zai-he-an-zhuang/install/docker)

### نصب دستی

[آموزش نصب دستی](https://Aqr-K.github.io/XrayR-doc/xrayr-xia-zai-he-an-zhuang/install/manual)

## فایل های پیکربندی و آموزش های با جرئیات

[آموزش مفصل](https://Aqr-K.github.io/XrayR-doc/)

## Thanks

- [Project X](https://github.com/XTLS/)
- [V2Fly](https://github.com/v2fly)
- [VNet-V2ray](https://github.com/ProxyPanel/VNet-V2ray)
- [Air-Universe](https://github.com/crossfw/Air-Universe)

## Licence

[Mozilla Public License Version 2.0](https://github.com/Aqr-K/XrayR/blob/master/LICENSE)

## Telgram

[بحث در مورد XrayR Backend](https://t.me/XrayR_project)

[کانال اعلان در مورد XrayR](https://t.me/XrayR_channel)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/Aqr-K/XrayR.svg)](https://starchart.cc/Aqr-K/XrayR)
