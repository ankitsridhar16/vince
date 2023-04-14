!function(){"use strict";var l=window.location,u=window.document,s=u.currentScript,d=s.getAttribute("data-api")||new URL(s.src).origin+"/api/event";function f(t){console.warn("Ignoring Event: "+t)}function t(t,e){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(l.hostname)||"file:"===l.protocol)return f("localhost");if(!(window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)){try{if("true"===window.localStorage.vince_ignore)return f("localStorage flag")}catch(t){}var r=s&&s.getAttribute("data-include"),n=s&&s.getAttribute("data-exclude");if("pageview"===t){var i=!r||r&&r.split(",").some(c),a=n&&n.split(",").some(c);if(!i||a)return f("exclusion rule")}var o={};o.n=t,o.u=e&&e.u?e.u:l.href,o.d=s.getAttribute("data-domain"),o.r=u.referrer||null,o.w=window.innerWidth,e&&e.meta&&(o.m=JSON.stringify(e.meta)),e&&e.props&&(o.p=e.props);var p=new XMLHttpRequest;p.open("POST",d,!0),p.setRequestHeader("Content-Type","text/plain"),p.send(JSON.stringify(o)),p.onreadystatechange=function(){4===p.readyState&&e&&e.callback&&e.callback()}}function c(t){return l.pathname.match(new RegExp("^"+t.trim().replace(/\*\*/g,".*").replace(/([^\.])\*/g,"$1[^\\s/]*")+"/?$"))}}var e=window.vince&&window.vince.q||[];window.vince=t;for(var r=0;r<e.length;r++)t.apply(this,e[r]);var i=1;function n(t){if("auxclick"!==t.type||t.button===i){var e,r=function(t){for(;t&&(void 0===t.tagName||(!(e=t)||!e.tagName||"a"!==e.tagName.toLowerCase())||!t.href);)t=t.parentNode;var e;return t}(t.target),n=r&&r.href&&r.href.split("?")[0];return(e=r)&&e.href&&e.host&&e.host!==l.host?a(t,r,{name:"Outbound Link: Click",props:{url:r.href}}):function(t){if(!t)return!1;var e=t.split(".").pop();return v.some(function(t){return t===e})}(n)?a(t,r,{name:"File Download",props:{url:n}}):void 0}}function a(t,e,r){var n=!1;function i(){n||(n=!0,window.location=e.href)}!function(t,e){if(!t.defaultPrevented){var r=!e.target||e.target.match(/^_(self|parent|top)$/i),n=!(t.ctrlKey||t.metaKey||t.shiftKey)&&"click"===t.type;return r&&n}}(t,e)?vince(r.name,{props:r.props}):(vince(r.name,{props:r.props,callback:i}),setTimeout(i,5e3),t.preventDefault())}u.addEventListener("click",n),u.addEventListener("auxclick",n);var o=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma"],p=s.getAttribute("file-types"),c=s.getAttribute("add-file-types"),v=p&&p.split(",")||c&&c.split(",").concat(o)||o}();