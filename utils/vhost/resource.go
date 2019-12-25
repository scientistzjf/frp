// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vhost

import (
	"bytes"
	"io/ioutil"
	"net/http"

	frpLog "github.com/fatedier/frp/utils/log"
	"github.com/fatedier/frp/utils/version"
)

var (
	NotFoundPagePath = ""
)

const (
	NotFound = `<!DOCTYPE HTML><html lang="en"><head><title>访问出错404</title><meta charset="utf-8" /><meta name="viewport" content="width=device-width,initial-scale=1,user-scalable=no" /><meta name="description" content="404页面" /><meta property="og:site_name" content="访问出错404" /><meta property="og:title" content="访问出错404" /><meta property="og:type" content="website" /><meta property="og:description" content="404页面" /><meta property="og:url" content="https://zjfe.top" />
<style>html,body,div,span,applet,object,iframe,h1,h2,h3,h4,h5,h6,p,blockquote,pre,a,abbr,acronym,address,big,cite,code,del,dfn,em,img,ins,kbd,q,s,samp,small,strike,strong,sub,sup,tt,var,b,u,i,center,dl,dt,dd,ol,ul,li,fieldset,form,label,legend,table,caption,tbody,tfoot,thead,tr,th,td,article,aside,canvas,details,embed,figure,figcaption,footer,header,hgroup,menu,nav,output,ruby,section,summary,time,mark,audio,video{margin:0;padding:0;border:0;font-size:100%;font:inherit;vertical-align:baseline;}article,aside,details,figcaption,figure,footer,header,hgroup,menu,nav,section{display:block;}body{line-height:1;}ol,ul{list-style:none;}blockquote,q{quotes:none;}blockquote:before,blockquote:after,q:before,q:after{content:'';content:none;}table{border-collapse:collapse;border-spacing:0;}body{-webkit-text-size-adjust:none}mark{background-color:transparent;color:inherit}input::-moz-focus-inner{border:0;padding:0}input[type="text"],input[type="email"],select,textarea{-moz-appearance:none;-webkit-appearance:none;-ms-appearance:none;appearance:none}*, *:before, *:after {box-sizing: border-box;}body {min-width: 320px;min-height: 100vh;line-height: 1.0;word-wrap: break-word;overflow-x: hidden;}u {text-decoration: underline;}strong {color: inherit;font-weight: bolder;}em {font-style: italic;}code {font-family: 'Lucida Console', 'Courier New', monospace;font-weight: normal;text-indent: 0;letter-spacing: 0;font-size: 0.9em;margin: 0 0.25em;padding: 0.25em 0.5em;background-color: rgba(144,144,144,0.25);border-radius: 0.25em;}mark {background-color: rgba(144,144,144,0.25);}a {-moz-transition: color 0.25s ease, background-color 0.25s ease, border-color 0.25s ease;-webkit-transition: color 0.25s ease, background-color 0.25s ease, border-color 0.25s ease;-ms-transition: color 0.25s ease, background-color 0.25s ease, border-color 0.25s ease;transition: color 0.25s ease, background-color 0.25s ease, border-color 0.25s ease;color: inherit;text-decoration: underline;}s {text-decoration: line-through;}body:before {content: '';display: block;background-attachment: scroll;position: fixed;top: 0;left: 0;width: 100vw;height: 100vh;z-index: 0;-moz-pointer-events: none;-webkit-pointer-events: none;-ms-pointer-events: none;pointer-events: none;-moz-transform: scale(1);-webkit-transform: scale(1);-ms-transform: scale(1);transform: scale(1);background-image: url('data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20viewBox%3D%220%200%20512%20512%22%20width%3D%22512%22%20height%3D%22512%22%20preserveAspectRatio%3D%22none%22%3E%3Cline%20x1%3D%22256%22%20y1%3D%220%22%20x2%3D%220%22%20y2%3D%22256%22%20stroke%3D%22rgba(167,184,119,0.502)%22%20stroke-width%3D%221.01px%22%20vector-effect%3D%22non-scaling-stroke%22%20stroke-linecap%3D%22square%22%20%2F%3E%3Cline%20x1%3D%22256%22%20y1%3D%22512%22%20x2%3D%22512%22%20y2%3D%22256%22%20stroke%3D%22rgba(167,184,119,0.502)%22%20stroke-width%3D%221.01px%22%20vector-effect%3D%22non-scaling-stroke%22%20stroke-linecap%3D%22square%22%20%2F%3E%3C%2Fsvg%3E'), linear-gradient(25deg, #CC960C 6%, #320987 61%);background-size: 272px, auto;background-position: center, 0% 0%;background-repeat: repeat, repeat;}body:after {display: block;content: '';position: fixed;top: 0;left: 0;width: 100%;height: 100%;background-color: #222836;z-index: 1;opacity: 0;visibility: hidden;-moz-transition: opacity 1s ease-in-out 1.125s, visibility 1s 1.125s;-webkit-transition: opacity 1s ease-in-out 1.125s, visibility 1s 1.125s;-ms-transition: opacity 1s ease-in-out 1.125s, visibility 1s 1.125s;transition: opacity 1s ease-in-out 1.125s, visibility 1s 1.125s;-moz-transform: scale(1);-webkit-transform: scale(1);-ms-transform: scale(1);transform: scale(1);}body.is-loading:after {opacity: 1;visibility: visible;}html {font-size: 18pt;}#wrapper {-webkit-overflow-scrolling: touch;display: flex;-moz-flex-direction: column;-webkit-flex-direction: column;-ms-flex-direction: column;flex-direction: column;-moz-align-items: center;-webkit-align-items: center;-ms-align-items: center;align-items: center;-moz-justify-content: center;-webkit-justify-content: center;-ms-justify-content: center;justify-content: center;min-height: 100vh;position: relative;z-index: 2;overflow: hidden;}#main {display: flex;position: relative;max-width: 100%;z-index: 1;-moz-align-items: center;-webkit-align-items: center;-ms-align-items: center;align-items: center;-moz-justify-content: center;-webkit-justify-content: center;-ms-justify-content: center;justify-content: center;-moz-flex-grow: 0;-webkit-flex-grow: 0;-ms-flex-grow: 0;flex-grow: 0;-moz-flex-shrink: 0;-webkit-flex-shrink: 0;-ms-flex-shrink: 0;flex-shrink: 0;text-align: left;-moz-transition: opacity 1s ease-in-out 0s;-webkit-transition: opacity 1s ease-in-out 0s;-ms-transition: opacity 1s ease-in-out 0s;transition: opacity 1s ease-in-out 0s;}#main > .inner {position: relative;z-index: 1;border-radius: inherit;padding: 3.625rem 2.875rem;max-width: 100%;width: 29rem;}#main > .inner > * {margin-top: 1.375rem;margin-bottom: 1.375rem;}#main > .inner > :first-child {margin-top: 0 !important;}#main > .inner > :last-child {margin-bottom: 0 !important;}#main > .inner > .full {margin-left: calc(-2.875rem);width: calc(100% + 5.75rem + 0.4725px);max-width: calc(100% + 5.75rem + 0.4725px);}#main > .inner > .full:first-child {margin-top: -3.625rem !important;border-top-left-radius: inherit;border-top-right-radius: inherit;}#main > .inner > .full:last-child {margin-bottom: -3.625rem !important;border-bottom-left-radius: inherit;border-bottom-right-radius: inherit;}#main > .inner > .full.screen {width: 100vw;max-width: 100vw;position: relative;border-radius: 0 !important;left: 50%;right: auto;margin-left: -50vw;}body.is-loading #main {opacity: 0;}body.is-instant #main, body.is-instant #main > .inner > *,body.is-instant #main > .inner > section > * {-moz-transition: none !important;-webkit-transition: none !important;-ms-transition: none !important;transition: none !important;}body.is-instant:after {display: none !important;-moz-transition: none !important;-webkit-transition: none !important;-ms-transition: none !important;transition: none !important;}h1 br + br, h2 br + br, h3 br + br, p br + br {display: block;content: ' ';}h1 .li, h2 .li, h3 .li, p .li {display: list-item;padding-left: 0.5em;margin: 0.75em 0 0 1em;}#text02 br + br {margin-top: 0.825rem;}#text02 {text-transform: uppercase;color: #FFFFFF;font-family: 'Montserrat', sans-serif;letter-spacing: 0.825rem;width: calc(100% + 0.825rem);font-size: 2.25em;line-height: 1.375;font-weight: 400;}#text01 br + br {margin-top: 1.05rem;}#text01 {color: #FFFFFF;font-family: 'Source Sans Pro', sans-serif;letter-spacing: 0.025rem;width: calc(100% + 0.025rem);font-size: 1em;line-height: 1.75;font-weight: 200;}#credits br + br {margin-top: 0.9rem;}#credits {color: #FFFFFF;font-family: 'Source Sans Pro', sans-serif;letter-spacing: 0.0125rem;width: calc(100% + 0.0125rem);font-size: 0.6em;line-height: 1.5;font-weight: 200;cursor: default;margin-top: 1rem !important;}#credits a {text-decoration: none !important;-moz-transition: opacity 0.25s ease;-webkit-transition: opacity 0.25s ease;-ms-transition: opacity 0.25s ease;transition: opacity 0.25s ease;opacity: 0.5;}#credits a:hover {text-decoration: none !important;opacity: 1;}.buttons {cursor: default;padding: 0;letter-spacing: 0;}.buttons li a {text-decoration: none;text-align: center;white-space: nowrap;max-width: 100%;-moz-align-items: center;-webkit-align-items: center;-ms-align-items: center;align-items: center;-moz-justify-content: center;-webkit-justify-content: center;-ms-justify-content: center;justify-content: center;vertical-align: middle;}.buttons.style1 {width: calc(100% + 1.75rem);margin-left: -0.875rem;}.buttons.style1 li {display: inline-block;vertical-align: middle;max-width: calc(100% - 1.75rem);margin: 0.875rem;}.buttons.style1 li a {display: flex;width: 8.875rem;height: 1.875rem;line-height: 1.875rem;vertical-align: middle;padding: 0 0.9375rem;text-transform: uppercase;font-size: 0.875em;font-family: 'IBM Plex Mono', monospace;letter-spacing: 0.325rem;padding-left: calc(0.325rem + 0.9375rem);font-weight: 600;border-radius: 1rem;}.buttons.style1 .button {background-color: #0C1845;color: #0E0D40;}.buttons.style1 .button:hover {background-color: #050724 !important;}#buttons01 .n01 {background-color: #307A71;color: #55F2B0;}#buttons01 .n01:hover {background-color: #F0DC26 !important;}.icons {cursor: default;padding: 0;letter-spacing: 0;}.icons li {display: inline-block;vertical-align: middle;position: relative;z-index: 1;}.icons li a {display: flex;-moz-align-items: center;-webkit-align-items: center;-ms-align-items: center;align-items: center;-moz-justify-content: center;-webkit-justify-content: center;-ms-justify-content: center;justify-content: center;}.icons li a svg {display: block;position: relative;-moz-transition: fill 0.25s ease;-webkit-transition: fill 0.25s ease;-ms-transition: fill 0.25s ease;transition: fill 0.25s ease;}.icons li a + svg {display: block;position: absolute;top: 0;left: 0;width: 100%;height: 100%;z-index: -1;-moz-pointer-events: none;-webkit-pointer-events: none;-ms-pointer-events: none;pointer-events: none;-moz-transition: fill 0.25s ease, stroke 0.25s ease;-webkit-transition: fill 0.25s ease, stroke 0.25s ease;-ms-transition: fill 0.25s ease, stroke 0.25s ease;transition: fill 0.25s ease, stroke 0.25s ease;}.icons li a .label {display: none;}#icons01 {font-size: 1.125em;width: calc(100% + 0.625rem);margin-left: -0.3125rem;}#icons01 li {margin: 0.3125rem;}#icons01 li a {border-radius: 100%;width: 2em;height: 2em;}#icons01 li a svg {width: 60%;height: 60%;}#icons01 a svg {fill: #FFFFFF;}#icons01 a {border: solid 1px transparent;}#icons01 a:hover {border-color: #1CFFE4 !important;}#icons01 a:hover svg {fill: #1CFFE4 !important;}@media (max-width: 1680px) {html {font-size: 13pt;}}@media (max-width: 1280px) {html {font-size: 13pt;}}@media (max-width: 980px) {html {font-size: 11pt;}}@media (max-width: 736px) {html {font-size: 11pt;}#main > .inner {padding: 3.5rem 2rem;}#main > .inner > * {margin-top: 1.375rem;margin-bottom: 1.375rem;}#main > .inner > .full {margin-left: calc(-2rem);width: calc(100% + 4rem + 0.4725px);max-width: calc(100% + 4rem + 0.4725px);}#main > .inner > .full:first-child {margin-top: -3.5rem !important;}#main > .inner > .full:last-child {margin-bottom: -3.5rem !important;}#main > .inner > .full.screen {margin-left: -50vw;}#text02 {letter-spacing: 0.721875rem;width: calc(100% + 0.721875rem);font-size: 2.25em;line-height: 1.375;}#text01 {letter-spacing: 0.021875rem;width: calc(100% + 0.021875rem);font-size: 1em;line-height: 1.75;}#credits {letter-spacing: 0.0109375rem;width: calc(100% + 0.0109375rem);font-size: 0.6em;line-height: 1.5;}.buttons.style1 li a {font-size: 0.875em;}#icons01 {font-size: 1.125em;width: calc(100% + 0.625rem);margin-left: -0.3125rem;}#icons01 li {margin: 0.3125rem;}}@media (max-width: 480px) {#main > .inner > * {margin-top: 1.203125rem;margin-bottom: 1.203125rem;}.buttons.style1 li a {width: 8.875rem;}}@media (max-width: 360px) {#main > .inner {padding: 2.625rem 1.5rem;}#main > .inner > * {margin-top: 1.03125rem;margin-bottom: 1.03125rem;}#main > .inner > .full {margin-left: calc(-1.5rem);width: calc(100% + 3rem + 0.4725px);max-width: calc(100% + 3rem + 0.4725px);}#main > .inner > .full:first-child {margin-top: -2.625rem !important;}#main > .inner > .full:last-child {margin-bottom: -2.625rem !important;}#main > .inner > .full.screen {margin-left: -50vw;}#text02 {font-size: 2em;}#text01 {font-size: 1em;}#credits {font-size: 0.6em;}.buttons.style1 {width: calc(100% + 1.3125rem);margin-left: -0.65625rem;}.buttons.style1 li {max-width: calc(100% - 1.3125rem);margin: 0.65625rem;}#icons01 {width: calc(100% + 0.46875rem);margin-left: -0.234375rem;}#icons01 li {margin: 0.234375rem;}}</style><noscript><style>body {overflow: auto !important;}body:after {display: none !important;}#main > .inner {opacity: 1.0 !important;}#main {opacity: 1.0 !important;-moz-transform: none !important;-webkit-transform: none !important;-ms-transform: none !important;transform: none !important;-moz-transition: none !important;-webkit-transition: none !important;-ms-transition: none !important;transition: none !important;-moz-filter: none !important;-webkit-filter: none !important;-ms-filter: none !important;filter: none !important;}</style></noscript></head><body class="is-loading"><svg xmlns="http://www.w3.org/2000/svg" version="1.1" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 40 40" display="none" width="0" height="0"><symbol id="icon-939" viewBox="0 0 40 40"><path d="M30.3,31.9L30.3,31.9c-1.4,1.5-2.9,2.4-4.6,3.1c-1.8,0.8-3.6,1.2-5.6,1.2s-3.8-0.4-5.6-1.1c-1.8-0.8-3.2-1.8-4.5-3 s-2.4-2.8-3-4.5c-0.5-1.2-0.8-2.4-1-3.3c0-0.3,0.2-0.5,0.9-0.6c0.6-0.1,0.9,0,1,0.4c0,0,0,0,0,0.1c0.2,1.1,0.5,2,0.8,2.7 c0.6,1.5,1.5,2.7,2.6,3.9c1.2,1.2,2.4,2.1,3.9,2.6c1.6,0.7,3.1,1,4.8,1s3.2-0.3,4.8-1c1.5-0.7,2.8-1.6,3.9-2.6l0.1-0.1 c0.1-0.1,0.3-0.2,0.4-0.1c0.1,0,0.3,0.2,0.6,0.4C30.4,31.3,30.4,31.7,30.3,31.9z M22,21.1l-1.2,1.2l1.1,1.1c0.3,0.3,0.2,0.5-0.1,0.9 c-0.2,0.2-0.4,0.3-0.6,0.3c-0.1,0-0.2-0.1-0.3-0.2l-1.1-1.1l-1.2,1.2c-0.1,0.1-0.1,0.1-0.3,0.1s-0.4-0.1-0.6-0.3l0,0 c-0.2-0.2-0.3-0.4-0.3-0.5s0-0.2,0.1-0.3l1.2-1.2l-1.2-1.2c-0.2-0.2-0.1-0.5,0.3-0.8c0.2-0.2,0.4-0.3,0.6-0.3c0.1,0,0.1,0,0.2,0.1 l1.2,1.2l1.2-1.2c0.2-0.2,0.5-0.1,0.9,0.2C22.1,20.7,22.2,20.9,22,21.1L22,21.1z M30.3,22.2c0,1.4-0.3,2.7-0.8,4 c-0.5,1.3-1.3,2.4-2.3,3.2c-1,1-2.1,1.7-3.2,2.3c-1.3,0.5-2.5,0.8-4,0.8c-1.4,0-2.7-0.3-4-0.8c-1.3-0.5-2.4-1.3-3.2-2.3 c-1-1-1.7-2.1-2.2-3.2c-0.2-0.4-0.3-0.6-0.3-0.7l0,0c-0.1-0.3,0.1-0.6,0.8-0.8c0.6-0.2,1-0.1,1.1,0.2c0.4,1.2,1,2.2,1.7,2.9l0,0v-6 c0-1.6,0.6-2.9,1.8-4c1.3-1.2,2.7-1.8,4.4-1.8c1.8,0,3.1,0.6,4.4,1.8c1.3,1.2,1.9,2.6,1.9,4.3c0,1.8-0.6,3.1-1.9,4.4 c-1.3,1.2-2.6,1.8-4.4,1.8c-0.7,0-1.4-0.1-2-0.3c-0.3-0.1-0.4-0.5-0.2-1.1c0.2-0.6,0.5-0.9,0.8-0.8l0.3,0.1c0.2,0,0.4,0.1,0.6,0.1 s0.4,0.1,0.5,0.1c1.2,0,2.3-0.4,3-1.3s1.3-1.9,1.3-3s-0.4-2.2-1.3-3c-0.8-0.8-1.9-1.3-3-1.3c-1.3,0-2.3,0.5-3.1,1.4 c-0.8,0.9-1.1,1.8-1.1,2.8v7.3c1.3,0.8,2.6,1.2,4.2,1.2c1.1,0,2.2-0.2,3.2-0.7c1.1-0.4,2-1.1,2.7-1.8c0.8-0.7,1.4-1.7,1.8-2.7 c0.4-1,0.7-2.1,0.7-3.2c0-2.4-0.8-4.3-2.4-6s-3.6-2.4-6-2.4s-4.3,0.8-6,2.4c-0.6,0.6-1.1,1.1-1.4,1.6l0,0c-0.1,0.1-0.2,0.2-0.2,0.3 c-0.1,0.1-0.2,0.1-0.4,0.2c-0.2,0.1-0.4,0-0.7-0.1c-0.3-0.1-0.5-0.2-0.7-0.3s-0.2-0.2-0.2-0.4V5.5c0-0.2,0.1-0.3,0.2-0.5 c0.1-0.2,0.2-0.2,0.4-0.2h15.4c0.4,0,0.5,0.3,0.5,1s-0.2,1-0.5,1H12.2v8.4l0,0c0.5-0.5,1.1-1,1.8-1.5c0.7-0.5,1.4-0.9,1.9-1.1 c1.3-0.5,2.6-0.8,4-0.8s2.7,0.3,4,0.8c1.3,0.5,2.4,1.3,3.2,2.3c1,1,1.7,2.1,2.3,3.2C30,19.4,30.3,20.8,30.3,22.2L30.3,22.2z M29.7,12c0.1,0.1,0.2,0.2,0.2,0.3s0,0.2-0.1,0.3c-0.1,0.1-0.2,0.2-0.3,0.4c-0.3,0.3-0.5,0.5-0.7,0.5c-0.1,0-0.2,0-0.3-0.1 c-1.3-1.1-2.5-1.9-3.6-2.4c-1.5-0.7-3.1-1-4.8-1c-1.6,0-3,0.3-4.6,0.9c-0.3,0.1-0.6-0.1-0.8-0.7c-0.1-0.3-0.2-0.5-0.1-0.7 c0-0.2,0.1-0.3,0.3-0.4c1.5-0.7,3.2-1,5.2-1s3.7,0.4,5.5,1.1C27.2,10,28.6,10.9,29.7,12L29.7,12z"/></symbol><symbol id="icon-906" viewBox="0 0 40 40"><path d="M27.2,4.7v4.9h-2.9c-1.1,0-1.8,0.2-2.1,0.6c-0.4,0.5-0.6,1.1-0.6,2v3.5H27l-0.7,5.4h-4.7v14H16v-14h-4.7v-5.4H16v-4.1 c0-2.3,0.6-4.1,1.9-5.3s2.9-1.9,5.2-1.9C24.8,4.4,26.2,4.5,27.2,4.7L27.2,4.7z"/></symbol><symbol id="icon-910" viewBox="0 0 40 40"><path d="M20,7c4.2,0,4.7,0,6.3,0.1c1.5,0.1,2.3,0.3,3,0.5C30,8,30.5,8.3,31.1,8.9c0.5,0.5,0.9,1.1,1.2,1.8c0.2,0.5,0.5,1.4,0.5,3 C33,15.3,33,15.8,33,20s0,4.7-0.1,6.3c-0.1,1.5-0.3,2.3-0.5,3c-0.3,0.7-0.6,1.2-1.2,1.8c-0.5,0.5-1.1,0.9-1.8,1.2 c-0.5,0.2-1.4,0.5-3,0.5C24.7,33,24.2,33,20,33s-4.7,0-6.3-0.1c-1.5-0.1-2.3-0.3-3-0.5C10,32,9.5,31.7,8.9,31.1 C8.4,30.6,8,30,7.7,29.3c-0.2-0.5-0.5-1.4-0.5-3C7,24.7,7,24.2,7,20s0-4.7,0.1-6.3c0.1-1.5,0.3-2.3,0.5-3C8,10,8.3,9.5,8.9,8.9 C9.4,8.4,10,8,10.7,7.7c0.5-0.2,1.4-0.5,3-0.5C15.3,7.1,15.8,7,20,7z M20,4.3c-4.3,0-4.8,0-6.5,0.1c-1.6,0-2.8,0.3-3.8,0.7 C8.7,5.5,7.8,6,6.9,6.9C6,7.8,5.5,8.7,5.1,9.7c-0.4,1-0.6,2.1-0.7,3.8c-0.1,1.7-0.1,2.2-0.1,6.5s0,4.8,0.1,6.5 c0,1.6,0.3,2.8,0.7,3.8c0.4,1,0.9,1.9,1.8,2.8c0.9,0.9,1.7,1.4,2.8,1.8c1,0.4,2.1,0.6,3.8,0.7c1.6,0.1,2.2,0.1,6.5,0.1 s4.8,0,6.5-0.1c1.6-0.1,2.9-0.3,3.8-0.7c1-0.4,1.9-0.9,2.8-1.8c0.9-0.9,1.4-1.7,1.8-2.8c0.4-1,0.6-2.1,0.7-3.8 c0.1-1.6,0.1-2.2,0.1-6.5s0-4.8-0.1-6.5c-0.1-1.6-0.3-2.9-0.7-3.8c-0.4-1-0.9-1.9-1.8-2.8c-0.9-0.9-1.7-1.4-2.8-1.8 c-1-0.4-2.1-0.6-3.8-0.7C24.8,4.3,24.3,4.3,20,4.3L20,4.3L20,4.3z"/><path d="M20,11.9c-4.5,0-8.1,3.7-8.1,8.1s3.7,8.1,8.1,8.1s8.1-3.7,8.1-8.1S24.5,11.9,20,11.9z M20,25.2c-2.9,0-5.2-2.3-5.2-5.2 s2.3-5.2,5.2-5.2s5.2,2.3,5.2,5.2S22.9,25.2,20,25.2z"/><path d="M30.6,11.6c0,1-0.8,1.9-1.9,1.9c-1,0-1.9-0.8-1.9-1.9s0.8-1.9,1.9-1.9C29.8,9.7,30.6,10.5,30.6,11.6z"/></symbol><symbol id="icon-905" viewBox="0 0 40 40"><path d="M36.3,10.2c-1,1.3-2.1,2.5-3.4,3.5c0,0.2,0,0.4,0,1c0,1.7-0.2,3.6-0.9,5.3c-0.6,1.7-1.2,3.5-2.4,5.1 c-1.1,1.5-2.3,3.1-3.7,4.3c-1.4,1.2-3.3,2.3-5.3,3c-2.1,0.8-4.2,1.2-6.6,1.2c-3.6,0-7-1-10.2-3c0.4,0,1.1,0.1,1.5,0.1 c3.1,0,5.9-1,8.2-2.9c-1.4,0-2.7-0.4-3.8-1.3c-1.2-1-1.9-2-2.2-3.3c0.4,0.1,1,0.1,1.2,0.1c0.6,0,1.2-0.1,1.7-0.2 c-1.4-0.3-2.7-1.1-3.7-2.3s-1.4-2.6-1.4-4.2v-0.1c1,0.6,2,0.9,3,0.9c-1-0.6-1.5-1.3-2.2-2.4c-0.6-1-0.9-2.1-0.9-3.3s0.3-2.3,1-3.4 c1.5,2.1,3.6,3.6,6,4.9s4.9,2,7.6,2.1c-0.1-0.6-0.1-1.1-0.1-1.4c0-1.8,0.8-3.5,2-4.7c1.2-1.2,2.9-2,4.7-2c2,0,3.6,0.8,4.8,2.1 c1.4-0.3,2.9-0.9,4.2-1.5c-0.4,1.5-1.4,2.7-2.9,3.6C33.8,11.2,35.1,10.9,36.3,10.2L36.3,10.2z"/></symbol></svg><div id="wrapper"><div id="main"><div class="inner"><h1 id="text02"><strong>404</strong></h1><p id="text01">不好意思，小站出错了</p><ul id="buttons01" class="style1 buttons"><li><a href="https://zjfe.top" class="button n01">三丰的网站</a></li></ul></div></div></div><script>(function() {var on = addEventListener,$ = function(q) { return document.querySelector(q) },$$ = function(q) { return document.querySelectorAll(q) },$body = document.body,$inner = $('.inner'),client = (function() {var o = {browser: 'other',browserVersion: 0,os: 'other',osVersion: 0,canUse: null},ua = navigator.userAgent,a, i;a = [['firefox', /Firefox\/([0-9\.]+)/],['edge', /Edge\/([0-9\.]+)/],['safari', /Version\/([0-9\.]+).+Safari/],['chrome', /Chrome\/([0-9\.]+)/],['ie', /Trident\/.+rv:([0-9]+)/]];for (i=0; i < a.length; i++) {if (ua.match(a[i][1])) {o.browser = a[i][0];o.browserVersion = parseFloat(RegExp.$1);break;}}a = [['ios', /([0-9_]+) like Mac OS X/, function(v) { return v.replace('_', '.').replace('_', ''); }],['ios', /CPU like Mac OS X/, function(v) { return 0 }],['android', /Android ([0-9\.]+)/, null],['mac', /Macintosh.+Mac OS X ([0-9_]+)/, function(v) { return v.replace('_', '.').replace('_', ''); }],['windows', /Windows NT ([0-9\.]+)/, null],['undefined', /Undefined/, null],];for (i=0; i < a.length; i++) {if (ua.match(a[i][1])) {o.os = a[i][0];o.osVersion = parseFloat( a[i][2] ? (a[i][2])(RegExp.$1) : RegExp.$1 );break;}}var _canUse = document.createElement('div');o.canUse = function(p) {var e = _canUse.style,up = p.charAt(0).toUpperCase() + p.slice(1);return (p in e|| ('Moz' + up) in e|| ('Webkit' + up) in e|| ('O' + up) in e|| ('ms' + up) in e);};return o;}()),trigger = function(t) {if (client.browser == 'ie') {var e = document.createEvent('Event');e.initEvent(t, false, true);dispatchEvent(e);} else dispatchEvent(new Event(t));},cssRules = function(selectorText) {var ss = document.styleSheets,a = [],f = function(s) {var r = s.cssRules,i;for (i=0; i < r.length; i++) {if (r[i] instanceof CSSMediaRule && matchMedia(r[i].conditionText).matches)(f)(r[i]); else if (r[i] instanceof CSSStyleRule && r[i].selectorText == selectorText)a.push(r[i]);}},x, i;for (i=0; i < ss.length; i++)f(ss[i]);return a;};var thisURL = function() {return window.location.href.replace(window.location.search, '').replace(/#$/, '');},getVar = function(name) {var a = window.location.search.substring(1).split('&'),b, k;for (k in a) {b = a[k].split('=');if (b[0] == name)return b[1];}return null;},cmd = function(cmd, values, handler) {var x, k, data;data = new FormData;data.append('cmd', cmd);for (k in values)data.append(k, values[k]);x = new XMLHttpRequest();x.open('POST', 'post/cmd');x.onreadystatechange = function() {var o;if (x.readyState != 4)return;if (x.status != 200)throw new Error('Failed server response (' + x.status + ')');try {o = JSON.parse(x.responseText);}catch (e) {throw new Error('Invalid server response');}if (!('result' in o)|| !('message' in o))throw new Error('Incomplete server response');if (o.result !== true)throw new Error(o.message);(handler)(o);};x.send(data);},redirectToStripeCheckout = function(options) {cmd('stripeCheckoutStart',options,function(response) {Stripe(options.key).redirectToCheckout({sessionId: response.sessionId}).then(function (result) {alert(result.error.message);});});},errors = {handle: function(handler) {window.onerror = function(message) {(handler)(message);return true;};},unhandle: function() {window.onerror = null;}},db = {open: function(objectStoreName, handler) {var request = indexedDB.open('carrd');request.onupgradeneeded = function(event) {event.target.result.createObjectStore(objectStoreName,{keyPath: 'id'});};request.onsuccess = function(event) {(handler)(event.target.result.transaction([objectStoreName],'readwrite').objectStore(objectStoreName));};},put: function(objectStore, values, handler) {var request = objectStore.put(values);request.onsuccess = function(event) {(handler)();};request.onerror = function(event) {throw new Error('db.put: error');};},get: function(objectStore, id, handler) {var request = objectStore.get(id);request.onsuccess = function(event) {if (!event.target.result)throw new Error('db.get: could not retrieve object with id "' + id + '"');(handler)(event.target.result);};request.onerror = function(event) {throw new Error('db.get: error');};},delete: function(objectStore, id, handler) {objectStore.delete(id).onsuccess = function(event) {(handler)(event.target.result);};},};on('load', function() {setTimeout(function() {$body.className = $body.className.replace(/\bis-loading\b/, 'is-playing');setTimeout(function() {$body.className = $body.className.replace(/\bis-playing\b/, 'is-ready');}, 2125);}, 100);});var style, sheet, rule;style = document.createElement('style');style.appendChild(document.createTextNode(''));document.head.appendChild(style);sheet = style.sheet;if (client.os == 'android') {(function() {sheet.insertRule('body::after { }', 0);rule = sheet.cssRules[0];var f = function() {rule.style.cssText = 'height: ' + (Math.max(screen.width, screen.height)) + 'px';};on('load', f);on('orientationchange', f);on('touchmove', f);})();} else if (client.os == 'ios') {if (client.osVersion <= 11)(function() {sheet.insertRule('body::after { }', 0);rule = sheet.cssRules[0];rule.style.cssText = '-webkit-transform: scale(1.0)';})();if (client.osVersion <= 11)(function() {sheet.insertRule('body.ios-focus-fix::before { }', 0);rule = sheet.cssRules[0];rule.style.cssText = 'height: calc(100% + 60px)';on('focus', function(event) {$body.classList.add('ios-focus-fix');}, true);on('blur', function(event) {$body.classList.remove('ios-focus-fix');}, true);})();} else if (client.browser == 'ie') {if (!('matches' in Element.prototype))Element.prototype.matches = (Element.prototype.msMatchesSelector || Element.prototype.webkitMatchesSelector);(function() {var a = cssRules('body::before'),r;if (a.length > 0) {r = a[0];if (r.style.width.match('calc')) {r.style.opacity = 0.9999;setTimeout(function() {r.style.opacity = 1;}, 100);} else {document.styleSheets[0].addRule('body::before', 'content: none !important;');$body.style.backgroundImage = r.style.backgroundImage.replace('url("images/', 'url("assets/images/');$body.style.backgroundPosition = r.style.backgroundPosition;$body.style.backgroundRepeat = r.style.backgroundRepeat;$body.style.backgroundColor = r.style.backgroundColor;$body.style.backgroundAttachment = 'fixed';$body.style.backgroundSize = r.style.backgroundSize;}}})();(function() {var t, f;f = function() {var mh, h, s, xx, x, i;x = $('#wrapper');x.style.height = 'auto';if (x.scrollHeight <= innerHeight)x.style.height = '100vh';xx = $$('.container.full');for (i=0; i < xx.length; i++) {x = xx[i];s = getComputedStyle(x);x.style.minHeight = '';x.style.height = '';mh = s.minHeight;x.style.minHeight = 0;x.style.height = '';h = s.height;if (mh == 0)continue;x.style.height = (h > mh ? 'auto' : mh);}};(f)();on('resize', function() {clearTimeout(t);t = setTimeout(f, 250);});on('load', f);})();} else if (client.browser == 'edge') {(function() {var xx = $$('.container > .inner > div:last-child'),x, y, i;for(i=0; i < xx.length; i++) {x = xx[i];y = getComputedStyle(x.parentNode);if (y.display != 'flex'&& y.display != 'inline-flex')continue;x.style.marginLeft = '-1px';}})();}if (!client.canUse('object-fit')) {(function() {var xx = $$('.image[data-position]'),x, w, c, i, src;for (i=0; i < xx.length; i++) {x = xx[i];c = x.firstElementChild;if (c.tagName != 'IMG') {w = c;c = c.firstElementChild;}if (c.parentNode.classList.contains('deferred')) {c.parentNode.classList.remove('deferred');src = c.getAttribute('data-src');c.removeAttribute('data-src');} else src = c.getAttribute('src');c.style['backgroundImage'] = 'url(\'' + src + '\')';c.style['backgroundSize'] = 'cover';c.style['backgroundPosition'] = x.dataset.position;c.style['backgroundRepeat'] = 'no-repeat';c.src = 'data:image/svg+xml;charset=utf8,' + escape('<svg xmlns="http://www.w3.org/2000/svg" width="1" height="1" viewBox="0 0 1 1"></svg>');if (x.classList.contains('full')&& (x.parentNode && x.parentNode.classList.contains('full'))&& (x.parentNode.parentNode && x.parentNode.parentNode.parentNode && x.parentNode.parentNode.parentNode.classList.contains('container'))&& x.parentNode.children.length == 1) {(function(x, w) {var p = x.parentNode.parentNode,f = function() {x.style['height'] = '0px';clearTimeout(t);t = setTimeout(function() {if (getComputedStyle(p).flexDirection == 'row') {if (w)w.style['height'] = '100%';x.style['height'] = (p.scrollHeight + 1) + 'px';} else {if (w)w.style['height'] = 'auto';x.style['height'] = 'auto';}}, 125);},t;on('resize', f);on('load', f);(f)();})(x, w);}}})();(function() {var xx = $$('.gallery img'),x, p, i, src;for (i=0;i < xx.length; i++) {x = xx[i];p = x.parentNode;if (p.classList.contains('deferred')) {p.classList.remove('deferred');src = x.getAttribute('data-src');} else src = x.getAttribute('src');p.style['backgroundImage'] = 'url(\'' + src + '\')';p.style['backgroundSize'] = 'cover';p.style['backgroundPosition'] = 'center';p.style['backgroundRepeat'] = 'no-repeat';x.style['opacity'] = '0';}})();}})();</script></body></html>
`
)

func getNotFoundPageContent() []byte {
	var (
		buf []byte
		err error
	)
	if NotFoundPagePath != "" {
		buf, err = ioutil.ReadFile(NotFoundPagePath)
		if err != nil {
			frpLog.Warn("read custom 404 page error: %v", err)
			buf = []byte(NotFound)
		}
	} else {
		buf = []byte(NotFound)
	}
	return buf
}

func notFoundResponse() *http.Response {
	header := make(http.Header)
	header.Set("server", "frp/"+version.Full())
	header.Set("Content-Type", "text/html")

	res := &http.Response{
		Status:     "Not Found",
		StatusCode: 404,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     header,
		Body:       ioutil.NopCloser(bytes.NewReader(getNotFoundPageContent())),
	}
	return res
}

func noAuthResponse() *http.Response {
	header := make(map[string][]string)
	header["WWW-Authenticate"] = []string{`Basic realm="Restricted"`}
	res := &http.Response{
		Status:     "401 Not authorized",
		StatusCode: 401,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     header,
	}
	return res
}
