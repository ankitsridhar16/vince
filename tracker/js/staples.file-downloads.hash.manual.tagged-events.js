(function(){"use strict";var n,s,i,a,c,l,d,h,r=window.location,e=window.document,t=e.currentScript,C=t.getAttribute("data-api")||x(t);function b(e){console.warn("Ignoring Event: "+e)}function x(e){return new URL(e.src).origin+"/api/event"}function u(n,s){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(r.hostname)||r.protocol==="file:")return b("localhost");if(window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)return;try{if(window.localStorage.vince_ignore==="true")return b("localStorage flag")}catch{}var i,o={};o.n=n,o.u=s&&s.u?s.u:r.href,o.d=t.getAttribute("data-domain"),o.r=e.referrer||null,o.w=window.innerWidth,s&&s.meta&&(o.m=JSON.stringify(s.meta)),s&&s.props&&(o.p=s.props),o.h=1,i=new XMLHttpRequest,i.open("POST",C,!0),i.setRequestHeader("Content-Type","text/plain"),i.send(JSON.stringify(o)),i.onreadystatechange=function(){i.readyState===4&&s&&s.callback&&s.callback()}}c=window.vince&&window.vince.q||[],window.vince=u;for(n=0;n<c.length;n++)u.apply(this,c[n]);function E(e){for(;e&&(typeof e.tagName=="undefined"||!p(e)||!e.href);)e=e.parentNode;return e}function p(e){return e&&e.tagName&&e.tagName.toLowerCase()==="a"}function O(e,t){if(e.defaultPrevented)return!1;var n=!t.target||t.target.match(/^_(self|parent|top)$/i),s=!(e.ctrlKey||e.metaKey||e.shiftKey)&&e.type==="click";return n&&s}s=1;function m(e){if(e.type==="auxclick"&&e.button!==s)return;var t=E(e.target),n=t&&t.href&&t.href.split("?")[0];if(g(t,0))return;if(y(n))return f(e,t,{name:"File Download",props:{url:n}})}function f(e,t,n){var s=!1;function o(){s||(s=!0,window.location=t.href)}O(e,t)?(vince(n.name,{props:n.props,callback:o}),setTimeout(o,5e3),e.preventDefault()):vince(n.name,{props:n.props})}e.addEventListener("click",m),e.addEventListener("auxclick",m),d=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma"],l=t.getAttribute("file-types"),a=t.getAttribute("add-file-types"),h=l&&l.split(",")||a&&a.split(",").concat(d)||d;function y(e){if(!e)return!1;var t=e.split(".").pop();return h.some(function(e){return e===t})}function j(e){var n,s,a,r,l,c=o(e)?e:e&&e.parentNode,t={name:null,props:{}},i=c&&c.classList;if(!i)return t;for(n=0;n<i.length;n++){if(l=i.item(n),s=l.match(/vince-event-(.+)=(.+)/),!s)continue;a=s[1],r=s[2].replace(/\+/g," "),a.toLowerCase()==="name"?t.name=r:t.props[a]=r}return t}function _(e){var n,s=e.target,t=j(s);if(!t.name)return;e.preventDefault(),n=!1;function o(){n||(n=!0,s.submit())}setTimeout(o,5e3),vince(t.name,{props:t.props,callback:o})}function w(e){return e&&e.tagName&&e.tagName.toLowerCase()==="form"}i=3;function v(e){if(e.type==="auxclick"&&e.button!==s)return;for(var n,a,r,t=e.target,c=0;c<=i;c++){if(!t)break;if(w(t))return;p(t)&&(a=t),o(t)&&(r=t),t=t.parentNode}r&&(n=j(r),a?(n.props.url=a.href,f(e,a,n)):vince(n.name,{props:n.props}))}function o(e){var t,n=e&&e.classList;if(n)for(t=0;t<n.length;t++)if(n.item(t).match(/vince-event-name=(.+)/))return!0;return!1}function g(e,t){return!(!e||t>i)&&(!!o(e)||g(e.parentNode,t+1))}e.addEventListener("submit",_),e.addEventListener("click",v),e.addEventListener("auxclick",v)})()