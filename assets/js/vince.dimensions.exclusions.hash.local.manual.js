!function(){"use strict";var s=window.location,g=window.document,d=g.currentScript,w=d.getAttribute("data-api")||new URL(d.src).origin+"/api/event";function v(e){console.warn("Ignoring Event: "+e)}function e(e,t){try{if("true"===window.localStorage.vince_ignore)return v("localStorage flag")}catch(e){}var n=d&&d.getAttribute("data-include"),r=d&&d.getAttribute("data-exclude");if("pageview"===e){var a=!n||n&&n.split(",").some(o),i=r&&r.split(",").some(o);if(!a||i)return v("exclusion rule")}function o(e){var t=s.pathname;return(t+=s.hash).match(new RegExp("^"+e.trim().replace(/\*\*/g,".*").replace(/([^\.])\*/g,"$1[^\\s/]*")+"/?$"))}var c={};c.n=e,c.u=t&&t.u?t.u:s.href,c.d=d.getAttribute("data-domain"),c.r=g.referrer||null,c.w=window.innerWidth,t&&t.meta&&(c.m=JSON.stringify(t.meta)),t&&t.props&&(c.p=t.props);var u=d.getAttributeNames().filter(function(e){return"event-"===e.substring(0,6)}),l=c.p||{};u.forEach(function(e){var t=e.replace("event-",""),n=d.getAttribute(e);l[t]=l[t]||n}),c.p=l,c.h=1;var p=new XMLHttpRequest;p.open("POST",w,!0),p.setRequestHeader("Content-Type","text/plain"),p.send(JSON.stringify(c)),p.onreadystatechange=function(){4===p.readyState&&t&&t.callback&&t.callback()}}var t=window.vince&&window.vince.q||[];window.vince=e;for(var n=0;n<t.length;n++)e.apply(this,t[n])}();