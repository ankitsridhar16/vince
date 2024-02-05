(function(){"use strict";var s,r,c,l,h,t=window.location,e=window.document,n=e.currentScript,C=n.getAttribute("data-api")||x(n);function i(e){console.warn("Ignoring Event: "+e)}function x(e){return new URL(e.src).origin+"/api/event"}function a(s,o){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(t.hostname)||t.protocol==="file:")return i("localhost");if(window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)return;try{if(window.localStorage.vince_ignore==="true")return i("localStorage flag")}catch{}var a,r,d,u,c=n&&n.getAttribute("data-include"),l=n&&n.getAttribute("data-exclude");if(s==="pageview"&&(d=!c||c&&c.split(",").some(h),u=l&&l.split(",").some(h),!d||u))return i("exclusion rule");function h(e){var n=t.pathname;return n+=t.hash,n.match(new RegExp("^"+e.trim().replace(/\*\*/g,".*").replace(/([^.])\*/g,"$1[^\\s/]*")+"/?$"))}a={},a.n=s,a.u=t.href,a.d=n.getAttribute("data-domain"),a.r=e.referrer||null,a.w=window.innerWidth,o&&o.meta&&(a.m=JSON.stringify(o.meta)),o&&o.props&&(a.p=o.props),a.h=1,r=new XMLHttpRequest,r.open("POST",C,!0),r.setRequestHeader("Content-Type","text/plain"),r.send(JSON.stringify(a)),r.onreadystatechange=function(){r.readyState===4&&o&&o.callback&&o.callback()}}r=window.vince&&window.vince.q||[],window.vince=a;for(s=0;s<r.length;s++)a.apply(this,r[s]);function o(){h=t.pathname,a("pageview")}window.addEventListener("hashchange",o);function O(){!h&&e.visibilityState==="visible"&&o()}e.visibilityState==="prerender"?e.addEventListener("visibilitychange",O):o();function b(e){for(;e&&(typeof e.tagName=="undefined"||!m(e)||!e.href);)e=e.parentNode;return e}function m(e){return e&&e.tagName&&e.tagName.toLowerCase()==="a"}function w(e,t){if(e.defaultPrevented)return!1;var n=!t.target||t.target.match(/^_(self|parent|top)$/i),s=!(e.ctrlKey||e.metaKey||e.shiftKey)&&e.type==="click";return n&&s}c=1;function g(e){if(e.type==="auxclick"&&e.button!==c)return;var t=b(e.target),n=t&&t.href&&t.href.split("?")[0];if(f(t,0))return;if(j(t))return v(e,t,{name:"Outbound Link: Click",props:{url:t.href}})}function v(e,t,n){var s=!1;function o(){s||(s=!0,window.location=t.href)}w(e,t)?(vince(n.name,{props:n.props,callback:o}),setTimeout(o,5e3),e.preventDefault()):vince(n.name,{props:n.props})}e.addEventListener("click",g),e.addEventListener("auxclick",g);function j(e){return e&&e.href&&e.host&&e.host!==t.host}function u(e){var n,s,i,a,c,r=d(e)?e:e&&e.parentNode,t={name:null,props:{}},o=r&&r.classList;if(!o)return t;for(n=0;n<o.length;n++){if(c=o.item(n),s=c.match(/vince-event-(.+)=(.+)/),!s)continue;i=s[1],a=s[2].replace(/\+/g," "),i.toLowerCase()==="name"?t.name=a:t.props[i]=a}return t}function y(e){var n,s=e.target,t=u(s);if(!t.name)return;e.preventDefault(),n=!1;function o(){n||(n=!0,s.submit())}setTimeout(o,5e3),vince(t.name,{props:t.props,callback:o})}function _(e){return e&&e.tagName&&e.tagName.toLowerCase()==="form"}l=3;function p(e){if(e.type==="auxclick"&&e.button!==c)return;for(var n,s,o,t=e.target,i=0;i<=l;i++){if(!t)break;if(_(t))return;m(t)&&(s=t),d(t)&&(o=t),t=t.parentNode}o&&(n=u(o),s?(n.props.url=s.href,v(e,s,n)):vince(n.name,{props:n.props}))}function d(e){var t,n=e&&e.classList;if(n)for(t=0;t<n.length;t++)if(n.item(t).match(/vince-event-name=(.+)/))return!0;return!1}function f(e,t){return!(!e||t>l)&&(!!d(e)||f(e.parentNode,t+1))}e.addEventListener("submit",y),e.addEventListener("click",p),e.addEventListener("auxclick",p)})()