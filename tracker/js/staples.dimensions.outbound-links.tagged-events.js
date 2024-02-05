(function(){"use strict";var s,i,a,c,d,u,v,t=window.location,e=window.document,n=e.currentScript,E=n.getAttribute("data-api")||_(n);function m(e){console.warn("Ignoring Event: "+e)}function _(e){return new URL(e.src).origin+"/api/event"}function l(s,o){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(t.hostname)||t.protocol==="file:")return m("localhost");if(window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)return;try{if(window.localStorage.vince_ignore==="true")return m("localStorage flag")}catch{}var a,r,c,i={};i.n=s,i.u=t.href,i.d=n.getAttribute("data-domain"),i.r=e.referrer||null,i.w=window.innerWidth,o&&o.meta&&(i.m=JSON.stringify(o.meta)),o&&o.props&&(i.p=o.props),c=n.getAttributeNames().filter(function(e){return e.substring(0,6)==="event-"}),r=i.p||{},c.forEach(function(e){var t=e.replace("event-",""),s=n.getAttribute(e);r[t]=r[t]||s}),i.p=r,a=new XMLHttpRequest,a.open("POST",E,!0),a.setRequestHeader("Content-Type","text/plain"),a.send(JSON.stringify(i)),a.onreadystatechange=function(){a.readyState===4&&o&&o.callback&&o.callback()}}d=window.vince&&window.vince.q||[],window.vince=l;for(s=0;s<d.length;s++)l.apply(this,d[s]);function o(){if(u===t.pathname)return;u=t.pathname,l("pageview")}i=window.history,i.pushState&&(v=i.pushState,i.pushState=function(){v.apply(this,arguments),o()},window.addEventListener("popstate",o));function C(){!u&&e.visibilityState==="visible"&&o()}e.visibilityState==="prerender"?e.addEventListener("visibilitychange",C):o();function k(e){for(;e&&(typeof e.tagName=="undefined"||!p(e)||!e.href);)e=e.parentNode;return e}function p(e){return e&&e.tagName&&e.tagName.toLowerCase()==="a"}function y(e,t){if(e.defaultPrevented)return!1;var n=!t.target||t.target.match(/^_(self|parent|top)$/i),s=!(e.ctrlKey||e.metaKey||e.shiftKey)&&e.type==="click";return n&&s}a=1;function b(e){if(e.type==="auxclick"&&e.button!==a)return;var t=k(e.target),n=t&&t.href&&t.href.split("?")[0];if(f(t,0))return;if(x(t))return j(e,t,{name:"Outbound Link: Click",props:{url:t.href}})}function j(e,t,n){var s=!1;function o(){s||(s=!0,window.location=t.href)}y(e,t)?(vince(n.name,{props:n.props,callback:o}),setTimeout(o,5e3),e.preventDefault()):vince(n.name,{props:n.props})}e.addEventListener("click",b),e.addEventListener("auxclick",b);function x(e){return e&&e.href&&e.host&&e.host!==t.host}function h(e){var n,s,i,a,l,c=r(e)?e:e&&e.parentNode,t={name:null,props:{}},o=c&&c.classList;if(!o)return t;for(n=0;n<o.length;n++){if(l=o.item(n),s=l.match(/vince-event-(.+)=(.+)/),!s)continue;i=s[1],a=s[2].replace(/\+/g," "),i.toLowerCase()==="name"?t.name=a:t.props[i]=a}return t}function w(e){var n,s=e.target,t=h(s);if(!t.name)return;e.preventDefault(),n=!1;function o(){n||(n=!0,s.submit())}setTimeout(o,5e3),vince(t.name,{props:t.props,callback:o})}function O(e){return e&&e.tagName&&e.tagName.toLowerCase()==="form"}c=3;function g(e){if(e.type==="auxclick"&&e.button!==a)return;for(var n,s,o,t=e.target,i=0;i<=c;i++){if(!t)break;if(O(t))return;p(t)&&(s=t),r(t)&&(o=t),t=t.parentNode}o&&(n=h(o),s?(n.props.url=s.href,j(e,s,n)):vince(n.name,{props:n.props}))}function r(e){var t,n=e&&e.classList;if(n)for(t=0;t<n.length;t++)if(n.item(t).match(/vince-event-name=(.+)/))return!0;return!1}function f(e,t){return!(!e||t>c)&&(!!r(e)||f(e.parentNode,t+1))}e.addEventListener("submit",w),e.addEventListener("click",g),e.addEventListener("auxclick",g)})()