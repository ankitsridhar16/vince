!function(){"use strict";var u=window.location,a=window.document,i=a.currentScript,o=i.getAttribute("data-api")||new URL(i.src).origin+"/api/event";function e(e,t){try{if("true"===window.localStorage.vince_ignore)return void console.warn("Ignoring Event: localStorage flag")}catch(e){}var n={};n.n=e,n.u=t&&t.u?t.u:u.href,n.d=i.getAttribute("data-domain"),n.r=a.referrer||null,n.w=window.innerWidth,t&&t.meta&&(n.m=JSON.stringify(t.meta)),t&&t.props&&(n.p=t.props);var r=new XMLHttpRequest;r.open("POST",o,!0),r.setRequestHeader("Content-Type","text/plain"),r.send(JSON.stringify(n)),r.onreadystatechange=function(){4===r.readyState&&t&&t.callback&&t.callback()}}var t=window.vince&&window.vince.q||[];window.vince=e;for(var n=0;n<t.length;n++)e.apply(this,t[n]);var p=1;function r(e){if("auxclick"!==e.type||e.button===p){var t,n,r,a,i,o=function(e){for(;e&&(void 0===e.tagName||(!(t=e)||!t.tagName||"a"!==t.tagName.toLowerCase())||!e.href);)e=e.parentNode;var t;return e}(e.target);o&&o.href&&o.href.split("?")[0];if((i=o)&&i.href&&i.host&&i.host!==u.host)return t=e,r={name:"Outbound Link: Click",props:{url:(n=o).href}},a=!1,void(!function(e,t){if(!e.defaultPrevented){var n=!t.target||t.target.match(/^_(self|parent|top)$/i),r=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type;return n&&r}}(t,n)?vince(r.name,{props:r.props}):(vince(r.name,{props:r.props,callback:c}),setTimeout(c,5e3),t.preventDefault()))}function c(){a||(a=!0,window.location=n.href)}}a.addEventListener("click",r),a.addEventListener("auxclick",r)}();