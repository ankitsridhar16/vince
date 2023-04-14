!function(){"use strict";var t,e,n,o=window.location,p=window.document,c=p.getElementById("vince"),s=c.getAttribute("data-api")||(t=c.src.split("/"),e=t[0],n=t[2],e+"//"+n+"/api/event");function i(t,e){try{if("true"===window.localStorage.vince_ignore)return void console.warn("Ignoring Event: localStorage flag")}catch(t){}var n={};n.n=t,n.u=o.href,n.d=c.getAttribute("data-domain"),n.r=p.referrer||null,n.w=window.innerWidth,e&&e.meta&&(n.m=JSON.stringify(e.meta)),e&&e.props&&(n.p=e.props);var i=c.getAttributeNames().filter(function(t){return"event-"===t.substring(0,6)}),r=n.p||{};i.forEach(function(t){var e=t.replace("event-",""),n=c.getAttribute(t);r[e]=r[e]||n}),n.p=r,n.h=1;var a=new XMLHttpRequest;a.open("POST",s,!0),a.setRequestHeader("Content-Type","text/plain"),a.send(JSON.stringify(n)),a.onreadystatechange=function(){4===a.readyState&&e&&e.callback&&e.callback()}}var r=window.vince&&window.vince.q||[];window.vince=i;for(var a,l=0;l<r.length;l++)i.apply(this,r[l]);function u(){a=o.pathname,i("pageview")}window.addEventListener("hashchange",u),"prerender"===p.visibilityState?p.addEventListener("visibilitychange",function(){a||"visible"!==p.visibilityState||u()}):u();var v=1;function d(t){if("auxclick"!==t.type||t.button===v){var e,n=function(t){for(;t&&(void 0===t.tagName||(!(e=t)||!e.tagName||"a"!==e.tagName.toLowerCase())||!t.href);)t=t.parentNode;var e;return t}(t.target),i=n&&n.href&&n.href.split("?")[0];return(e=n)&&e.href&&e.host&&e.host!==o.host?f(t,n,{name:"Outbound Link: Click",props:{url:n.href}}):function(t){if(!t)return!1;var e=t.split(".").pop();return h.some(function(t){return t===e})}(i)?f(t,n,{name:"File Download",props:{url:i}}):void 0}}function f(t,e,n){var i=!1;function r(){i||(i=!0,window.location=e.href)}!function(t,e){if(!t.defaultPrevented){var n=!e.target||e.target.match(/^_(self|parent|top)$/i),i=!(t.ctrlKey||t.metaKey||t.shiftKey)&&"click"===t.type;return n&&i}}(t,e)?vince(n.name,{props:n.props}):(vince(n.name,{props:n.props,callback:r}),setTimeout(r,5e3),t.preventDefault())}p.addEventListener("click",d),p.addEventListener("auxclick",d);var g=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma"],w=c.getAttribute("file-types"),m=c.getAttribute("add-file-types"),h=w&&w.split(",")||m&&m.split(",").concat(g)||g}();