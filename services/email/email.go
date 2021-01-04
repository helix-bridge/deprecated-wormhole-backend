package email

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/darwinia-network/link/util"
	"github.com/jordan-wright/email"
)

func SendToSubscribe(to string) {
	username := util.GetEnv("SMTP_USERNAME", "")
	password := util.GetEnv("SMTP_PASSWORD", "")
	e := email.NewEmail()
	e.From = "Darwinia Network <noreply@darwinia.network>"
	e.To = []string{to}
	e.Subject = "You've Subscribed to Darwinia Network!"
	e.HTML = []byte(SubscribeEmail)
	if os.Getenv("GIN_MODE") == "release" {
		if err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", username, password, "smtp.gmail.com")); err != nil {
			fmt.Println("Send mail error ...........", err)
		}
	}
}

var SubscribeEmail = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
        "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html style="width:100%;font-family:helvetica, 'helvetica neue', arial, verdana, sans-serif;-webkit-text-size-adjust:100%;-ms-text-size-adjust:100%;padding:0;Margin:0;">
<head>
    <meta charset="UTF-8">
    <meta content="width=device-width, initial-scale=1" name="viewport">
    <meta name="x-apple-disable-message-reformatting">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta content="telephone=no" name="format-detection">
    <title>Darwinia Network</title>
    <!--[if (mso 16)]>
    <style type="text/css">    a {
        text-decoration: none;
    }    </style>    <![endif]-->
    <!--[if gte mso 9]>
    <style>sup {
        font-size: 100% !important;
    }</style><![endif]-->
    <style type="text/css">


        #outlook a {
            padding: 0;
        }

        a[x-apple-data-detectors] {
            color: inherit !important;
            text-decoration: none !important;
            font-size: inherit !important;
            font-family: inherit !important;
            font-weight: inherit !important;
            line-height: inherit !important;
        }

        .es-desk-hidden {
            display: none;
            float: left;
            overflow: hidden;
            width: 0;
            max-height: 0;
            line-height: 0;
            mso-hide: all;
        }

        .es-button-border:hover a.es-button {
            background: #ffffff !important;
            border-color: #ffffff !important;
        }

        .es-button-border:hover {
            background: #ffffff !important;
            border-style: solid solid solid solid !important;
            border-color: #3d5ca3 #3d5ca3 #3d5ca3 #3d5ca3 !important;
        }
    </style>
</head>
<body style="width:100%;font-family:helvetica, 'helvetica neue', arial, verdana, sans-serif;-webkit-text-size-adjust:100%;-ms-text-size-adjust:100%;padding:0;Margin:0;">
<div class="es-wrapper-color" style="background-color:#FAFAFA;">
    <!--[if gte mso 9]>
    <v:background xmlns:v="urn:schemas-microsoft-com:vml" fill="t">
        <v:fill type="tile" color="#fafafa"></v:fill>
    </v:background><![endif]-->
    <table class="es-wrapper" width="100%" cellspacing="0" cellpadding="0"
           style="mso-table-lspace:0pt;min-width: 600px; mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;padding-top:87px;padding-bottom:87px;Margin:0;width:100%;height:100%;background-repeat:repeat;background-position:center top;background:linear-gradient(315deg,rgba(254,56,118,1) 0%,rgba(124,48,221,1) 71%,rgba(58,48,221,1) 100%);">
        <tr style="border-collapse:collapse;">
            <td valign="top" style="padding-top:87px;padding-bottom:87px;margin: 0;">
                <table class="es-content" cellspacing="0" cellpadding="0" align="center"
                       style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;table-layout:fixed !important;width:100%;">
                    <tr style="border-collapse:collapse;">
                        <td class="es-info-area" style="padding:0;Margin:0;"
                            align="center">
                            <table class="es-content-body"
                                   style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;"
                                   width="600" cellspacing="0" cellpadding="0" align="center">
                                <tr style="border-collapse:collapse;">
                                    <td style="Margin:0;padding-bottom:5px;padding-top:20px;background-position:left top;"
                                        align="left">
                                        <table width="100%" cellspacing="0" cellpadding="0"
                                               style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;">
                                            <tr style="border-collapse:collapse;">
                                                <td width="560" valign="top" align="left" style="padding:0;Margin:0;">
                                                    <table width="100%" cellspacing="0" cellpadding="0"
                                                           style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;">
                                                        <tr style="border-collapse:collapse;">
                                                            <td class="es-infoblock" align="left"
                                                                style="padding:0;Margin:0;padding-bottom:20px;font-size:12px;color:#fff;">
                                                                <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-size:30px;font-family:helvetica, 'helvetica neue', arial, verdana, sans-serif;color:#fff;">
                                                                    Darwinia Network Newsletter </p>
                                                            </td>
                                                            <td align="right" style="padding-bottom: 12px;">
                                                                <a target="_blank"
                                                                   href="var://@unsubscribe()"
                                                                   style="-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:helvetica, 'helvetica neue', arial, verdana, sans-serif;font-size:12px;text-decoration:none;color:#fff;display: inline-block;padding: 5px 28px;opacity:0.6;border:1px solid rgba(255,255,255,1);">Unsubscribe</a>
                                                            </td>
                                                        </tr>
                                                    </table>
                                                </td>
                                            </tr>
                                        </table>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                </table>
                <table class="es-content" cellspacing="0" cellpadding="0" align="center"
                       style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;table-layout:fixed !important;width:100%;">
                    <tr style="border-collapse:collapse;">
                        <td style="padding:0;Margin:0;"  align="center">
                            <table class="es-content-body"
                                   style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;background-color:#FFFFFF;"
                                   width="600" cellspacing="0" cellpadding="0" bgcolor="#ffffff" align="center">
                                <tr style="border-collapse:collapse;">
                                    <td style="Margin:0;padding-top:50px;padding-bottom:15px;padding-left:45px;padding-right:45px;border-radius:10px 10px 0 0px;background-position:left top;"
                                         align="left">
                                        <table width="100%" cellspacing="0" cellpadding="0"
                                               style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;">
                                            <tr style="border-collapse:collapse;">
                                                <td width="560" valign="top" align="center" style="padding:0;Margin:0;">
                                                    <table width="100%" cellspacing="0" cellpadding="0"
                                                           style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;">
                                                        <tr style="border-collapse:collapse;">
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-top:10px;padding-bottom: 20px;border-bottom: 2px solid #EFEFEF;font-size: 24px;text-transform: uppercase;">Subscription Confirmed</td>
                                                        </tr>
                                                    </table>
                                                </td>
                                            </tr>
                                        </table>
                                    </td>
                                </tr>
                                <tr style="border-collapse:collapse;">
                                    <td style="padding:0;Margin:0;padding-left:20px;padding-right:20px;padding-top:0px;background-color:transparent;background-position:left top;"
                                        bgcolor="transparent" align="left">
                                        <table width="100%" cellspacing="0" cellpadding="0"
                                               style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;">
                                            <tr style="border-collapse:collapse;">
                                                <td width="560" valign="top" align="center" style="padding:0;Margin:0;">
                                                    <table style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;background-position:left top;"
                                                           width="100%" cellspacing="0" cellpadding="0">
                                                        <tr style="border-collapse:collapse;">
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-top:15px;padding-left:26px;padding-right:26px;padding-bottom: 110px;">
                                                                <p style="Margin:0;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-size:14px;font-family:helvetica, 'helvetica neue', arial, verdana, sans-serif;line-height:24px;color:#000;">
                                                                    Congratulations! Your subscription to our list has been confirmed.<br/>

                                                                    Thank you for subscription!<br/>

                                                                </p>
                                                                <p style="Margin:0;margin-top: 60px;-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-size:14px;font-family:helvetica, 'helvetica neue', arial, verdana, sans-serif;line-height:24px;color:#000;">
                                                                    Darwinia Network<br/>
                                                                    Polkadot Parachain For User-facing Blockchain Application Developers<br/>
                                                                </p>
                                                            </td>
                                                        </tr>

                                                    </table>
                                                </td>
                                            </tr>
                                        </table>
                                    </td>
                                </tr>


                            </table>
                        </td>
                    </tr>
                </table>
                <table class="es-footer" cellspacing="0" cellpadding="0" align="center"
                       style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;table-layout:fixed !important;width:100%;background-color:transparent;background-repeat:repeat;background-position:center top;">
                    <tr style="border-collapse:collapse;">
                        <td style="padding:0;Margin:0;" align="center">
                            <table class="es-footer-body" width="600" cellspacing="0" cellpadding="0"
                                   align="center"
                                   style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;background-color:transparent;">
                                <tr style="border-collapse:collapse;">
                                    <td style="Margin:0;padding-top:60px;padding-left:45px;padding-right:0px;padding-bottom:50px;border-radius:0px 0px 0px 0px;border-bottom:2px solid #352b60;background-color:#201550;background-position:left top;"
                                        bgcolor="#201550" align="left">
                                        <table width="100%" cellspacing="0" cellpadding="0"
                                               style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;">
                                            <tr style="border-collapse:collapse;">
                                                <td width="560" valign="top" align="center" style="padding:0;Margin:0;font-size: 14px;">
                                                    <table width="100%" cellspacing="0" cellpadding="0"
                                                           style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;">
                                                        <tr style="border-collapse:collapse;color: #fff">
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-top:5px;padding-bottom:35px;width:33%">
                                                                GENERAL
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-top:5px;padding-bottom:35px;width:34%">
                                                                TECHNOLOGY
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-top:5px;padding-bottom:35px;width:33%">
                                                                COMMUNITY
                                                            </td>
                                                        </tr>

                                                        <tr style="border-collapse:collapse;color: #7B70AE">
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://darwinia.network/#top">About</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://docs.darwinia.network/docs/en/crab-home">Darwinia Crab Network</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://docs.darwinia.network/docs/en/wiki-rfc-index">RFCS</a>
                                                            </td>
                                                        </tr>

                                                        <tr style="border-collapse:collapse;color: #7B70AE">
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://docs.darwinia.network/">WIKI</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://darwinia.network/Darwinia_Genepaper_EN.pdf">Genepaper</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="http://darwinia.network/ambassador">Ambassador Program</a>
                                                            </td>
                                                        </tr>

                                                        <tr style="border-collapse:collapse;color: #7B70AE">
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://www.itering.io/about">Careers</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://github.com/darwinia-network/darwinia">Github</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="http://darwinia.network/community">Get involved</a>
                                                            </td>
                                                        </tr>

                                                        <tr style="border-collapse:collapse;color: #7B70AE">
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="mailto:hello@darwinia.network">Cooperation</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://telemetry.polkadot.io/#list/Polkadot">Telemetry</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="http://darwinia.network/faq">FAQ</a>
                                                            </td>
                                                        </tr>

                                                        <tr style="border-collapse:collapse;color: #7B70AE">
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="http://darwinia.network/brand">Brand</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:5px;padding-bottom:5px;">
                                                            </td>
                                                        </tr>
                                                    </table>
                                                </td>
                                            </tr>
                                        </table>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                </table>

                <table class="es-footer" cellspacing="0" cellpadding="0" align="center"
                       style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;table-layout:fixed !important;width:100%;background-color:transparent;background-repeat:repeat;background-position:center top;">
                    <tr style="border-collapse:collapse;">
                        <td style="padding:0;Margin:0;" align="center">
                            <table class="es-footer-body" width="600" cellspacing="0" cellpadding="0"
                                   align="center"
                                   style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;background-color:transparent;">
                                <tr style="border-collapse:collapse;">
                                    <td style="Margin:0;padding-top:24px;padding-left:52px;padding-right:52px;padding-bottom:24px;border-radius:0px 0px 0px 0px;background-color:#201550;background-position:left top;"
                                        bgcolor="#201550" align="left">
                                        <table width="100%" cellspacing="0" cellpadding="0"
                                               style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;">
                                            <tr style="border-collapse:collapse;">
                                                <td width="560" valign="top" align="center" style="padding:0;Margin:0;font-size: 14px;">
                                                    <table width="100%" cellspacing="0" cellpadding="0"
                                                           style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;">
                                                        <tr style="border-collapse:collapse;color: #fff">
                                                            <td align="left" style="padding:0;Margin:0;color: #7B70AE;">
                                                                Copyright@2020 Darwinia Network
                                                            </td>
                                                            <td align="right" style="padding:0;Margin:0;">
                                                                <a style="margin-left: 10px;color: #7B70AE;text-decoration: none;" target="_blank" href="https://t.me/DarwiniaNetwork">
                                                                    <img style="width: 25px; height: 25px" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAENElEQVRoQ+1YTYhbVRT+Tm4y0RlJqgtR6EJdVBftwi4KFkHEYmmlM5MgERe1dFGFUYa80HZhXkhezSDOZCZpsWg6P4ILF0XzXkdNQamo0MGKP3QoVFCLtFVadVFb/zpJ7pGkZniGZPLy8zpB3lu+e75zv+98957zEoqGc4z/wUOOkB5z0XGkxwyB44jjiE0VcI6WTYVtO63jSNulswnoOGJTYdtO2zuOEH0H0BsEvgzwODPuaEXVagspADCEi7La5PBHRFT5SaFGcq+yxPO9L4RwDsC0RwzMJVJbf64lrIZzrzPwXE8KIaIiGPMgZF+aGv6wWv16ZFVFP8nMm3tMCP1ALp4RwjurTTxxyUwuHs/78FfhNm186Kfq+6NHWSwuGFcZ3L/qQv6t/nsgV1b4Bj/QNJK1pGKKHpLMMepDKDkePFtdjynGBslysRUR5diuXnYCnQcwI/pcs+Yq/8eFfe/fVSosvQZgIzzuLcmJHd+a11XF2M0s51ZDSAlEecGuLPkHj9er/nK1I8YzkjkD4BcQbRmbGr5QSziq6IfBPHLThBBwEaAZ0eed1ca3X1xp4/j+/Nri0vUjAG8jwmm3GHi8Xrcq54gq+ikwb7JVCBEkA8cFXNn1Dw3lQyEqNdtQjeT2QFKKwT4QPnP7aJumBa7Uw2WzX3jOn71wDWBvs7y165buCAG/E1HG1SemtVcGy/eg6aPunb+XS6VpMD9WCSY6sUb0D+1Lbf2jETgent9YRPHLpsnrBDQVUnYBbs8DtZey0WbMTDHFeIEJL4N5oKIBmL/zvltCo6Pbr69EMhbWn5XgrC1CbhChhPDjYKMjUd04HsmtK0maZfDD1XdE9JbwrdmlaY8WmxFUFeMIs9zTLK7eelNHTKACEU4Qu94RQhiJyR2/VtfKQ+z0wjEFJA+AceuyCCAr/IGRlTqZmZSq6F8x84N2CzHnLxHhExC9TUxnJDhV22mIXBPJ9PB+q6QOHcp7L5/7+xoAj1WMOa4VRyznJxfU5FRwzDIAQFzRNxWZT7WCsVUIEe1NpgOTrRKKKcaIZHm4VdzyMe72v/Fuj/fu2o9DK+RUJTfHjN1WYju97Jb2cPvXeTVt/ZKlYFOQquiLzLyhVZwtjhDoz2QmUJkd5o7W7AsgkXi3v3ClcBWA6AkhIPw4lg6urZKJRnIRSNpJHjyZnAh834hkNGJshpQn2xVRxnW7a50ZywQrx0NVDIVZTlXIEX4TwK4D6eCxemRjEWNUSnmwZ4QQ8GkyE3wkFtHDUnK6llh5tgif/8XaKa8q+pvMvLNnhIDoczAWAA43IlUWKzzep8ydTQ3r3zD4/t4RYpEJES5B0NPJVODjTr54bR2IFrXcuDpEXzNwD5hvbwV3U+ZIp4TaxXe7a7XLo2OcI6TjEnY5geNIlwvacTrHkY5L2OUEjiNdLmjH6f4BIQ7x8oqsRLYAAAAASUVORK5CYII="/>
                                                                </a>
                                                                <a style="margin-left: 10px;color: #7B70AE;text-decoration: none;" target="_blank" href="https://twitter.com/DarwiniaNetwork">
                                                                    <img style="width: 25px; height: 25px" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAELElEQVRoQ+1YX2hbVRz+fufetZG1cTrdZGjZ2MMQ5oMMReeLhTlkbE0zZcp8kPnigy9JjR0l0TRLOofY3rAHYVNRkPniSLoSOup0PijioIigmyKz3ZgwqmhdrWn+3Ht/koTSrrbNOac2jHHzer/v9/2+33fuPeeEoqEM4zb4kWfkFkvRS+QWCwReIl4iqzQBb2mt0mC1y3qJaI9ulYheIisdLIH+AthkoGWltSr8xiZCuC4gUj72ZaPpp69XGohFhrbAtQ/CpQiD182ZoisEFBm8TcaokhEi+hSMSQY/J1P8JgzR5/41Inj4rcDfi3Hjh4fa7JL9IYHGCPyRYaz9sezkx8C8VkZLyYgh6Nl7NjfnJsaLn4B5n4xAFUN0yfSv25FItBdkOCdOjK659tO1HDPvJkKJGU31eEpGBJnbk1bHxXj8C9OZmjzJjEP1BKo+IJ5KpTs/k8HG4yycqcFDIGxg8GUwfSCTipIRE01bE+m9Y7MNvRHOBBymNMCbl2qSgF+TVrCNiJRuorHXslth81fMuE9mAEpGDMPYfaQ/cG5+4YHw13f8SRM97KIb4Ob/itLZvnRwj0wzs5hYd+ZBLtHwcgNaWE/JCIje6bOCryzW1NGezPp8EfvBdICZ2wEYtWVFmVQ6+IySkVB2lME7VDhKRgiYNnz0UOJY8MpyIr2RkQ2OM7OLyd3CwEzfwP4BlaZioexVBrepcBSMEIN4HMBvLc1Ne3qO7Z1UEVLBRsPZaZkXfH5NBSOAeeddrZi+0eYSWpNvd15QaU4We/z4cPPEWEHqM61vRJiPJwY6vpFtSgfXG8ltL9ul71W5SomA6GSfFXxZVUQFH+3KdMFFvwqn+lFR/IPOESTak1bnl6pCsvhYOHuOmXfJ4mdxqkYq3v8g0POyO7VKQ5XzllOyx5khVHg6iVTrE8Fl0HsmxLsJKzCqKroUPhYaTDHcqE49jUQqRijkg3m67G+5kUi0T+sIL+TEu89sckruxZuP8vKV9YyAfmltEg8vdSSXl59DRsPZIaUT9QIRLSO15UW5u7HxQJe1c0an8fmc18ODB112T62kjraRmih9RzBeTaU7zus2UTnOlO38JYDX69bQftkXChLhKkDJlBV8X6WZeHzYb08VRsB4TIW3GHaFiVRL/iBIvJm0Oj9WaeZoz/C9/xSKOTA/qsJbCitppHIpYqpGCMozoXL6PU9MZ3T2k1gk+yRsnGLwpv/DhPTSIiK7tm9Qf8IKXNYVj3cP3++UC0cAvKiz6S2nK5nIbAkqEnFWEJ0WfvNsb+++vIypaNfgTmJ+CeAXmOGT4ahiFI3Mla8tMR4B6FsB/EwkfnfARcEwmLERhAfA/AiAJ1QvSaompJeWTuFGc7QTaXSj9fQ8I/Um1OjnXiKNnng9PS+RehNq9HMvkUZPvJ7ebZPIv2+b3fLh+Yn7AAAAAElFTkSuQmCC"/>
                                                                </a>
                                                                <a style="margin-left: 10px;color: #7B70AE;text-decoration: none;" target="_blank" href="https://darwinianetwork.medium.com/">
                                                                    <img style="width: 25px; height: 25px" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAE2UlEQVRoQ+1Za2gcVRQ+587mpWlKfeCPSKCFPiQoIjXSgtFaLVpjd2exYK0/oqLBiu3O2ii6u1m3OyslaXeWIFihIlUoNpWZTar54WuNj6BUECTB2lqqxaZKG9PEmsZk5x6Z1oTsZmZ2E+lmqLv/ds433znf/e49984MBnwqwRXww6IQh7lYdMRhhkDRkaIjl2kEilPrMg3snGn/P44wBrsQhO5cQ6UT3wlEdblw2XGGwiYiOmt7H/JHieBxO0xORxDwuIAVdRHl/j/siFqkpKgTV2cjBAE+lxPeu+zuCfiTq4F4NxAs/E9CLt6MkHJVLVoXiaxJW5GFw8T0Ye0IASzNVwxDtjmqePZb4YO+rnsI9S4gujoXZ05HpgiQ7YkpnmfsCIP+ZBNxvidX0ktxHLxhSVn11q3r/zbDh6SuWoL0t0RQng9f/kKM1Ay3yHHxdWtXUuXp4aGTAHB9ruQMMB5NiM9buiFp24gokYtnMj47IYADckKstp3TktoCBJFcBWCpsEJudf9ohQv5NR/npOTimZMQQByKKeI1duSvvqReO3oBTxLQVVY4BPxMTohr7HgKIqSjg4T+fqBIBLlZMUG/+hpxeNaqUKPlRhX3u+FwymXVQAoiJBxOVernh9bJu72m7Tbse3+JjuPHiIBli0HEs0LV0uqKsdMVoxOjt+2Iu1Omi70QU8sQkh4Z+iimeFdZL1b1IBE8nB03Ntho3NsckpJbkPCXHQn3B/MrZHjoT4asPqp4vjArJCxpdWmibzJjSFjiWi63PXQs4NO+E4AFHSEEEA/FFHGD9Yam9hBA/VRnQfhUVrxrw1LnyjTphwUQGpwhBJCwlGrlVu8PFou+gTgcmoox9kgs7jkQkLQ3gOhpBwkBQIS3ZMX7hJkQIsKQlOwnoJsA4Ixr4bIbK/86VXJOHz0NRAucJmRcKBEWR1rdA6auSNqTRLSXMWyNxsUXg//+N7COEmIUNFmkmZD29u6y30+MnXCBUB9R3D8FfdrXBHSHI4UAwrCrqrwmElk/YtpKpeSdRncLScmbOfHvJzGOc+RSYaw5lvDssjt2BP1aO3F6ztlCEE7VrKhZ3NS0csJ0XwmnytMj5waAaJGzhRjVITTGFO8+i536Mc7pnemx+ZtagIOxhHhdXOqtGKTfRk0K7pMV8RZEnPGpIujL3Bzne7F/ElPEe40igpLaRwS1Mw6FjD0oxz0ZLyuCL3Qup3H9SDZ2Xhy5eGotg9sjO8WfjYKMDkTEewmgcnqBiNgjK+Ld06+F/Gob57DdCUImGLK12QfEFkl164QaAOH0Ipngqovu3nDYuBYO95Wmh4/+OuPxF+GCC+HWSNx71HRNbVdX8TR+lc1t1RXzetRlwJ6KJjx7zUgCUvJlIB7LdAXekxXvRuNawK9uBA4dGXGA84jQEFW8PXbtOuTXGonoTbPnmhlTOuenN4T2mOLdZrs/+LT9BLRpEoMIHFy4TG4Tjwd86ocAcN9UDHCEGD4Qi3t67TgnYwFf52YA3eiEgh3e1hFEGEMUVjMGuh0JT8MCDvqXWWtlH5DwNoH+ceb0YM0glBzMR8QURh9/BYAa5yxkVsnmGZzXGpnnGvNKXxSS1zAVEFR0pICDnVeqoiN5DVMBQUVHCjjYeaW6chwJStqZvCQ7HPQP33kQ/6mh3c4AAAAASUVORK5CYII="/>
                                                                </a>
                                                                <a style="margin-left: 10px;color: #7B70AE;text-decoration: none;" target="_blank" href="https://github.com/darwinia-network">
                                                                    <img style="width: 25px; height: 25px" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAEQUlEQVRoQ+2ZX4hUZRjGn/fMZDMYy0qFFGRFBbFKsheFihdJF2l/nD8X0h+hEMLqouaMtlkzw3F2zmiZejZDyLDsJrIl5syyJRJ2WSZJZW4UXWRBKS0ZrtTS7s6cJ2a3yW1z9Pu+3VlE5lw/z/u8v+/9vnPmnJFMqkRcBpe0QC6xKbYmcokNBK2JtCbSpBUw3loiCAAJSIab1JtWWWMQCPyil0w6G/vuCKp8OADXg5ynky6QMwK8hTDfv8oKDZwdCw6BvEunRl1rDCIWsu7OZLFeyHH89soQPBE+AuIYRb60wN8InCFFBGiHyDUEO0EuFpH9IURTeW/l7/UaGdvfDfLp2QWR0CrXix2cGuo4A3Py+UWjF2qmkSaXLq0LArw5qyAhhB7o7ol9aBLayJNL+48HAfeZ1DTfWpBn3J7EayahjTzZlJ8j2G1S0xgEAqfoJY1CG4LY5W1k8NwsgsivYYl2TD6oJuFTPcUX++YPDwffALxat57RRCxL7MLORI9umIo+k/K7AL6sop2s0QYRwWioLXJtPn/fWd0wFf3WTR/M+3NkbFD3QWsC8onrJZerNGWqydilwyCW6Pi1QSDW60Uv/pROiK42a5ffIIMndHzaIJZgS8FLZnRCdLU5u7w1YLBJx6cNImK94nrxLp0QXW3W9reT3KDjMwDBPtdLrtMJ0dVmbf9tko/p+LRBAPmq2JPo1AnR1WZSpeMAFun4tEFq7yHh0NzrNm+/d1AnSFXrdPVdXxkNfgYoqp6aThtkvLhIruglXJ0gVW027XczYE5VX9cZgdReiKKIdmR6Vp7SDbyQPpMu3yABBwi26dY1ApkYinwUarvtwYu9e6g2tGvXgSsHfxg5SPBuVc9knTHIP1usL9zWvjafX/GHSXjd4zgH2ipDI/sBrjKtowEiIyL4mEQHwJvOBcqPYlnPFnas7hcRrQ/iJCW3sRxDFa8SXGAKoXfYBafCnLN84bL7f/r6U383gfX/CRb5HuA7EOuQuyN2uBFUby9Dx474S6Uq91CwFuSt0wEwPOxy2hJZUfDixzO2XwYZ+18TIv1FL7G6UXOOQ6s65H9G4M6ZADAEqZ1y/DL/5sgtp09Yc6sY/oLEjf8WE6lEGV1wsbtZLt3XSVaPkrBmCkbjjJyLDImV7Pbi/uYN/bdXgrG9pCwDGIjIe66XeFSluaxdGiCxUEWrojECEciRghdfWj8He/YcveLkyYjo3IqzduldEg+pNKmiMQIZL2xh7+IliSfXrJGqStBUTTZVeonA8ybe83nMQSYeiidA9ool341/Bw4Hn7vbkt+qNJdN+VsIvqCiVdFMC2RqgIiVdr24pxLcAmmwSq2JnG9hdL53mf5cb7RtZ3QiAKoQ+QtgRSikYPy3l4Akav8sTFxCCIQREhGV86SimWkQlcymaFogTVnWaRRtTWQai9cUa2siTVnWaRS9bCbyN++dCwFuqzC6AAAAAElFTkSuQmCC"/>
                                                                </a>
                                                            </td>
                                                        </tr>
                                                    </table>
                                                </td>
                                            </tr>
                                        </table>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</div>
</body>
</html>`
