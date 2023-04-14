!function(){"use strict";var o=window.location,c=window.document,p=c.currentScript,s=p.getAttribute("data-api")||new URL(p.src).origin+"/api/event";function l(t){console.warn("Ignoring Event: "+t)}function t(t,e){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(o.hostname)||"file:"===o.protocol)return l("localhost");if(!(window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)){try{if("true"===window.localStorage.vince_ignore)return l("localStorage flag")}catch(t){}var n={};n.n=t,n.u=o.href,n.d=p.getAttribute("data-domain"),n.r=c.referrer||null,n.w=window.innerWidth,e&&e.meta&&(n.m=JSON.stringify(e.meta)),e&&e.props&&(n.p=e.props);var i=p.getAttributeNames().filter(function(t){return"event-"===t.substring(0,6)}),r=n.p||{};i.forEach(function(t){var e=t.replace("event-",""),n=p.getAttribute(t);r[e]=r[e]||n}),n.p=r,n.h=1;var a=new XMLHttpRequest;a.open("POST",s,!0),a.setRequestHeader("Content-Type","text/plain"),a.send(JSON.stringify(n)),a.onreadystatechange=function(){4===a.readyState&&e&&e.callback&&e.callback()}}}var e=window.vince&&window.vince.q||[];window.vince=t;for(var n,i=0;i<e.length;i++)t.apply(this,e[i]);function r(){n=o.pathname,t("pageview")}window.addEventListener("hashchange",r),"prerender"===c.visibilityState?c.addEventListener("visibilitychange",function(){n||"visible"!==c.visibilityState||r()}):r();var a=1;function u(t){if("auxclick"!==t.type||t.button===a){var e,n=function(t){for(;t&&(void 0===t.tagName||(!(e=t)||!e.tagName||"a"!==e.tagName.toLowerCase())||!t.href);)t=t.parentNode;var e;return t}(t.target),i=n&&n.href&&n.href.split("?")[0];return(e=n)&&e.href&&e.host&&e.host!==o.host?d(t,n,{name:"Outbound Link: Click",props:{url:n.href}}):function(t){if(!t)return!1;var e=t.split(".").pop();return g.some(function(t){return t===e})}(i)?d(t,n,{name:"File Download",props:{url:i}}):void 0}}function d(t,e,n){var i=!1;function r(){i||(i=!0,window.location=e.href)}!function(t,e){if(!t.defaultPrevented){var n=!e.target||e.target.match(/^_(self|parent|top)$/i),i=!(t.ctrlKey||t.metaKey||t.shiftKey)&&"click"===t.type;return n&&i}}(t,e)?vince(n.name,{props:n.props}):(vince(n.name,{props:n.props,callback:r}),setTimeout(r,5e3),t.preventDefault())}c.addEventListener("click",u),c.addEventListener("auxclick",u);var f=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma"],v=p.getAttribute("file-types"),w=p.getAttribute("add-file-types"),g=v&&v.split(",")||w&&w.split(",").concat(f)||f}();