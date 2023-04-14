!function(){"use strict";var l=window.location,f=window.document,v=f.currentScript,d=v.getAttribute("data-api")||new URL(v.src).origin+"/api/event";function m(e){console.warn("Ignoring Event: "+e)}function e(e,t){try{if("true"===window.localStorage.vince_ignore)return m("localStorage flag")}catch(e){}var n=v&&v.getAttribute("data-include"),r=v&&v.getAttribute("data-exclude");if("pageview"===e){var i=!n||n&&n.split(",").some(o),a=r&&r.split(",").some(o);if(!i||a)return m("exclusion rule")}function o(e){return l.pathname.match(new RegExp("^"+e.trim().replace(/\*\*/g,".*").replace(/([^\.])\*/g,"$1[^\\s/]*")+"/?$"))}var p={};p.n=e,p.u=l.href,p.d=v.getAttribute("data-domain"),p.r=f.referrer||null,p.w=window.innerWidth,t&&t.meta&&(p.m=JSON.stringify(t.meta)),t&&t.props&&(p.p=t.props);var c=v.getAttributeNames().filter(function(e){return"event-"===e.substring(0,6)}),u=p.p||{};c.forEach(function(e){var t=e.replace("event-",""),n=v.getAttribute(e);u[t]=u[t]||n}),p.p=u;var s=new XMLHttpRequest;s.open("POST",d,!0),s.setRequestHeader("Content-Type","text/plain"),s.send(JSON.stringify(p)),s.onreadystatechange=function(){4===s.readyState&&t&&t.callback&&t.callback()}}var t=window.vince&&window.vince.q||[];window.vince=e;for(var n,r=0;r<t.length;r++)e.apply(this,t[r]);function i(){n!==l.pathname&&(n=l.pathname,e("pageview"))}var a,o=window.history;function p(e){return e&&e.tagName&&"a"===e.tagName.toLowerCase()}o.pushState&&(a=o.pushState,o.pushState=function(){a.apply(this,arguments),i()},window.addEventListener("popstate",i)),"prerender"===f.visibilityState?f.addEventListener("visibilitychange",function(){n||"visible"!==f.visibilityState||i()}):i();var c=1;function u(e){if("auxclick"!==e.type||e.button===c){var t,n=function(e){for(;e&&(void 0===e.tagName||!p(e)||!e.href);)e=e.parentNode;return e}(e.target),r=n&&n.href&&n.href.split("?")[0];if(!function e(t,n){if(!t||k<n)return!1;if(x(t))return!0;return e(t.parentNode,n+1)}(n,0))return(t=n)&&t.href&&t.host&&t.host!==l.host?s(e,n,{name:"Outbound Link: Click",props:{url:n.href}}):function(e){if(!e)return!1;var t=e.split(".").pop();return b.some(function(e){return e===t})}(r)?s(e,n,{name:"File Download",props:{url:r}}):void 0}}function s(e,t,n){var r=!1;function i(){r||(r=!0,window.location=t.href)}!function(e,t){if(!e.defaultPrevented){var n=!t.target||t.target.match(/^_(self|parent|top)$/i),r=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type;return n&&r}}(e,t)?vince(n.name,{props:n.props}):(vince(n.name,{props:n.props,callback:i}),setTimeout(i,5e3),e.preventDefault())}f.addEventListener("click",u),f.addEventListener("auxclick",u);var g=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma"],w=v.getAttribute("file-types"),h=v.getAttribute("add-file-types"),b=w&&w.split(",")||h&&h.split(",").concat(g)||g;function y(e){var t=x(e)?e:e&&e.parentNode,n={name:null,props:{}},r=t&&t.classList;if(!r)return n;for(var i=0;i<r.length;i++){var a,o,p=r.item(i).match(/vince-event-(.+)=(.+)/);p&&(a=p[1],o=p[2].replace(/\+/g," "),"name"===a.toLowerCase()?n.name=o:n.props[a]=o)}return n}var k=3;function L(e){if("auxclick"!==e.type||e.button===c){for(var t,n,r,i,a=e.target,o=0;o<=k&&a;o++){if((r=a)&&r.tagName&&"form"===r.tagName.toLowerCase())return;p(a)&&(t=a),x(a)&&(n=a),a=a.parentNode}n&&(i=y(n),t?(i.props.url=t.href,s(e,t,i)):vince(i.name,{props:i.props}))}}function x(e){var t=e&&e.classList;if(t)for(var n=0;n<t.length;n++)if(t.item(n).match(/vince-event-name=(.+)/))return 1}f.addEventListener("submit",function(e){var t,n=e.target,r=y(n);function i(){t||(t=!0,n.submit())}r.name&&(e.preventDefault(),t=!1,setTimeout(i,5e3),vince(r.name,{props:r.props,callback:i}))}),f.addEventListener("click",L),f.addEventListener("auxclick",L)}();