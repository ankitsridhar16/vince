!function(){"use strict";var u=window.location,s=window.document,l=s.currentScript,f=l.getAttribute("data-api")||new URL(l.src).origin+"/api/event";function v(e){console.warn("Ignoring Event: "+e)}function e(e,t){try{if("true"===window.localStorage.vince_ignore)return v("localStorage flag")}catch(e){}var n=l&&l.getAttribute("data-include"),r=l&&l.getAttribute("data-exclude");if("pageview"===e){var a=!n||n&&n.split(",").some(o),i=r&&r.split(",").some(o);if(!a||i)return v("exclusion rule")}function o(e){return u.pathname.match(new RegExp("^"+e.trim().replace(/\*\*/g,".*").replace(/([^\.])\*/g,"$1[^\\s/]*")+"/?$"))}var c={};c.n=e,c.u=u.href,c.d=l.getAttribute("data-domain"),c.r=s.referrer||null,c.w=window.innerWidth,t&&t.meta&&(c.m=JSON.stringify(t.meta)),t&&t.props&&(c.p=t.props);var p=new XMLHttpRequest;p.open("POST",f,!0),p.setRequestHeader("Content-Type","text/plain"),p.send(JSON.stringify(c)),p.onreadystatechange=function(){4===p.readyState&&t&&t.callback&&t.callback()}}var t=window.vince&&window.vince.q||[];window.vince=e;for(var n,r=0;r<t.length;r++)e.apply(this,t[r]);function a(){n!==u.pathname&&(n=u.pathname,e("pageview"))}var i,o=window.history;function c(e){return e&&e.tagName&&"a"===e.tagName.toLowerCase()}o.pushState&&(i=o.pushState,o.pushState=function(){i.apply(this,arguments),a()},window.addEventListener("popstate",a)),"prerender"===s.visibilityState?s.addEventListener("visibilitychange",function(){n||"visible"!==s.visibilityState||a()}):a();var p=1;function d(e){var t;"auxclick"===e.type&&e.button!==p||((t=function(e){for(;e&&(void 0===e.tagName||!c(e)||!e.href);)e=e.parentNode;return e}(e.target))&&t.href&&t.href.split("?")[0],function e(t,n){if(!t||w<n)return!1;if(y(t))return!0;return e(t.parentNode,n+1)}(t,0))}function m(e,t,n){var r=!1;function a(){r||(r=!0,window.location=t.href)}!function(e,t){if(!e.defaultPrevented){var n=!t.target||t.target.match(/^_(self|parent|top)$/i),r=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type;return n&&r}}(e,t)?vince(n.name,{props:n.props}):(vince(n.name,{props:n.props,callback:a}),setTimeout(a,5e3),e.preventDefault())}function g(e){var t=y(e)?e:e&&e.parentNode,n={name:null,props:{}},r=t&&t.classList;if(!r)return n;for(var a=0;a<r.length;a++){var i,o,c=r.item(a).match(/vince-event-(.+)=(.+)/);c&&(i=c[1],o=c[2].replace(/\+/g," "),"name"===i.toLowerCase()?n.name=o:n.props[i]=o)}return n}s.addEventListener("click",d),s.addEventListener("auxclick",d);var w=3;function h(e){if("auxclick"!==e.type||e.button===p){for(var t,n,r,a,i=e.target,o=0;o<=w&&i;o++){if((r=i)&&r.tagName&&"form"===r.tagName.toLowerCase())return;c(i)&&(t=i),y(i)&&(n=i),i=i.parentNode}n&&(a=g(n),t?(a.props.url=t.href,m(e,t,a)):vince(a.name,{props:a.props}))}}function y(e){var t=e&&e.classList;if(t)for(var n=0;n<t.length;n++)if(t.item(n).match(/vince-event-name=(.+)/))return 1}s.addEventListener("submit",function(e){var t,n=e.target,r=g(n);function a(){t||(t=!0,n.submit())}r.name&&(e.preventDefault(),t=!1,setTimeout(a,5e3),vince(r.name,{props:r.props,callback:a}))}),s.addEventListener("click",h),s.addEventListener("auxclick",h)}();