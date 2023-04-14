!function(){"use strict";var l=window.location,v=window.document,d=v.currentScript,f=d.getAttribute("data-api")||new URL(d.src).origin+"/api/event";function g(e){console.warn("Ignoring Event: "+e)}function e(e,t){try{if("true"===window.localStorage.vince_ignore)return g("localStorage flag")}catch(e){}var n=d&&d.getAttribute("data-include"),i=d&&d.getAttribute("data-exclude");if("pageview"===e){var r=!n||n&&n.split(",").some(o),a=i&&i.split(",").some(o);if(!r||a)return g("exclusion rule")}function o(e){var t=l.pathname;return(t+=l.hash).match(new RegExp("^"+e.trim().replace(/\*\*/g,".*").replace(/([^\.])\*/g,"$1[^\\s/]*")+"/?$"))}var c={};c.n=e,c.u=l.href,c.d=d.getAttribute("data-domain"),c.r=v.referrer||null,c.w=window.innerWidth,t&&t.meta&&(c.m=JSON.stringify(t.meta)),t&&t.props&&(c.p=t.props);var p=d.getAttributeNames().filter(function(e){return"event-"===e.substring(0,6)}),u=c.p||{};p.forEach(function(e){var t=e.replace("event-",""),n=d.getAttribute(e);u[t]=u[t]||n}),c.p=u,c.h=1;var s=new XMLHttpRequest;s.open("POST",f,!0),s.setRequestHeader("Content-Type","text/plain"),s.send(JSON.stringify(c)),s.onreadystatechange=function(){4===s.readyState&&t&&t.callback&&t.callback()}}var t=window.vince&&window.vince.q||[];window.vince=e;for(var n,i=0;i<t.length;i++)e.apply(this,t[i]);function r(){n=l.pathname,e("pageview")}window.addEventListener("hashchange",r),"prerender"===v.visibilityState?v.addEventListener("visibilitychange",function(){n||"visible"!==v.visibilityState||r()}):r();var p=1;function a(e){if("auxclick"!==e.type||e.button===p){var t,n,i,r,a=function(e){for(;e&&(void 0===e.tagName||(!(t=e)||!t.tagName||"a"!==t.tagName.toLowerCase())||!e.href);)e=e.parentNode;var t;return e}(e.target),o=a&&a.href&&a.href.split("?")[0];if(function(e){if(!e)return!1;var t=e.split(".").pop();return s.some(function(e){return e===t})}(o))return r=!(i={name:"File Download",props:{url:o}}),void(!function(e,t){if(!e.defaultPrevented){var n=!t.target||t.target.match(/^_(self|parent|top)$/i),i=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type;return n&&i}}(t=e,n=a)?vince(i.name,{props:i.props}):(vince(i.name,{props:i.props,callback:c}),setTimeout(c,5e3),t.preventDefault()))}function c(){r||(r=!0,window.location=n.href)}}v.addEventListener("click",a),v.addEventListener("auxclick",a);var o=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma"],c=d.getAttribute("file-types"),u=d.getAttribute("add-file-types"),s=c&&c.split(",")||u&&u.split(",").concat(o)||o}();