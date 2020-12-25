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

var SubscribeEmail = `Subscribe<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
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
                                                                   style="-webkit-text-size-adjust:none;-ms-text-size-adjust:none;mso-line-height-rule:exactly;font-family:helvetica, 'helvetica neue', arial, verdana, sans-serif;font-size:12px;text-decoration:none;color:#fff;display: inline-block;padding: 5px 20px;border-radius:4px;opacity:0.6;border:1px solid rgba(255,255,255,1);">Unsubscribe</a>
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
                                                                style="padding:0;Margin:0;padding-top:10px;padding-bottom: 20px;border-bottom: 1px solid #ccc;font-size: 24px;text-transform: uppercase;">Subscription Confirmed</td>
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
                                    <td style="Margin:0;padding-top:50px;padding-left:52px;padding-right:0px;padding-bottom:30px;border-radius:0px 0px 0px 0px;background-color:#201550;background-position:left top;"
                                        bgcolor="#201550" align="left">
                                        <table width="100%" cellspacing="0" cellpadding="0"
                                               style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;">
                                            <tr style="border-collapse:collapse;">
                                                <td width="560" valign="top" align="center" style="padding:0;Margin:0;font-size: 14px;">
                                                    <table width="100%" cellspacing="0" cellpadding="0"
                                                           style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;">
                                                        <tr style="border-collapse:collapse;color: #fff">
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-top:9px;padding-bottom:35px;width:33%">
                                                                General
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-top:9px;padding-bottom:35px;width:34%">
                                                                Technology
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-top:9px;padding-bottom:35px;width:33%">
                                                                Community
                                                            </td>
                                                        </tr>

                                                        <tr style="border-collapse:collapse;color: #7B70AE">
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://mp.weixin.qq.com/s/vSWn8Wz2C_f3_mxOiV-54g">Staking</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://github.com/darwinia-network/darwinia/wiki/How-To-Join-Darwinia-POC-1-Testnet---Trilobita">Testnet</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://github.com/darwinia-network/rfcs">RFCS</a>
                                                            </td>
                                                        </tr>

                                                        <tr style="border-collapse:collapse;color: #7B70AE">
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://darwinia.network/faq">FAQ</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://telemetry.polkadot.io/#list/Darwinia%20POC-1%20Testnet">Telemetry</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://github.com/darwinia-network">GITHUB</a>
                                                            </td>
                                                        </tr>

                                                        <tr style="border-collapse:collapse;color: #7B70AE">
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="mailto:hello@darwinia.network">Contact</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://evolution.l2me.com/darwinia/Darwinia_Genepaper_EN.pdf">Genepaper</a>
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://medium.com/@DarwiniaNetwork">Medium</a>
                                                            </td>
                                                        </tr>

                                                        <tr style="border-collapse:collapse;color: #7B70AE">
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">

                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">

                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://t.me/DarwiniaNetwork">Telegram</a>
                                                            </td>
                                                        </tr>

                                                        <tr style="border-collapse:collapse;color: #7B70AE">
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">
                                                            </td>
                                                            <td align="left"
                                                                style="padding:0;Margin:0;padding-left: 1px;padding-top:9px;padding-bottom:5px;">
                                                                <a style="color: #7B70AE;text-decoration: none;" target="_blank" href="https://www.reddit.com/r/DarwiniaNetwork/">Reddit</a>
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
                                    <td style="Margin:0;padding-top:24px;padding-left:52px;padding-right:52px;padding-bottom:24px;border-radius:0px 0px 0px 0px;background-color:#000000;background-position:left top;"
                                        bgcolor="#201550" align="left">
                                        <table width="100%" cellspacing="0" cellpadding="0"
                                               style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;">
                                            <tr style="border-collapse:collapse;">
                                                <td width="560" valign="top" align="center" style="padding:0;Margin:0;font-size: 14px;">
                                                    <table width="100%" cellspacing="0" cellpadding="0"
                                                           style="mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:collapse;border-spacing:0px;">
                                                        <tr style="border-collapse:collapse;color: #fff">
                                                            <td align="center" style="padding:0;Margin:0;">
                                                                <img style="display: block" width="161" height="18" title="Darwinia" alt="Darwinia" src="https://darwinia.network/static/image/slide-logo.png"/>
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
