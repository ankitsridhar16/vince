(function(){"use strict";var s,i,c,n=window.location,t=window.document,e=t.currentScript,l=e.getAttribute("data-api")||d(e);function r(e){console.warn("Ignoring Event: "+e)}function d(e){return new URL(e.src).origin+"/api/event"}function o(s,o){try{if(window.localStorage.vince_ignore==="true")return r("localStorage flag")}catch{}var i,a,u,h,c=e&&e.getAttribute("data-include"),d=e&&e.getAttribute("data-exclude");if(s==="pageview"&&(u=!c||c&&c.split(",").some(m),h=d&&d.split(",").some(m),!u||h))return r("exclusion rule");function m(e){var t=n.pathname;return t+=n.hash,t.match(new RegExp("^"+e.trim().replace(/\*\*/g,".*").replace(/([^.])\*/g,"$1[^\\s/]*")+"/?$"))}i={},i.n=s,i.u=n.href,i.d=e.getAttribute("data-domain"),i.r=t.referrer||null,i.w=window.innerWidth,o&&o.meta&&(i.m=JSON.stringify(o.meta)),o&&o.props&&(i.p=o.props),i.h=1,a=new XMLHttpRequest,a.open("POST",l,!0),a.setRequestHeader("Content-Type","text/plain"),a.send(JSON.stringify(i)),a.onreadystatechange=function(){a.readyState===4&&o&&o.callback&&o.callback()}}i=window.vince&&window.vince.q||[],window.vince=o;for(s=0;s<i.length;s++)o.apply(this,i[s]);function a(){c=n.pathname,o("pageview")}window.addEventListener("hashchange",a);function u(){!c&&t.visibilityState==="visible"&&a()}t.visibilityState==="prerender"?t.addEventListener("visibilitychange",u):a()})()