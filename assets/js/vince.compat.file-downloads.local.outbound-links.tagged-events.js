!function(){"use strict";var e,t,n,i=window.location,a=window.document,o=a.getElementById("vince"),p=o.getAttribute("data-api")||(e=o.src.split("/"),t=e[0],n=e[2],t+"//"+n+"/api/event");function r(e,t){try{if("true"===window.localStorage.vince_ignore)return void console.warn("Ignoring Event: localStorage flag")}catch(e){}var n={};n.n=e,n.u=i.href,n.d=o.getAttribute("data-domain"),n.r=a.referrer||null,n.w=window.innerWidth,t&&t.meta&&(n.m=JSON.stringify(t.meta)),t&&t.props&&(n.p=t.props);var r=new XMLHttpRequest;r.open("POST",p,!0),r.setRequestHeader("Content-Type","text/plain"),r.send(JSON.stringify(n)),r.onreadystatechange=function(){4===r.readyState&&t&&t.callback&&t.callback()}}var c=window.vince&&window.vince.q||[];window.vince=r;for(var s,u=0;u<c.length;u++)r.apply(this,c[u]);function l(){s!==i.pathname&&(s=i.pathname,r("pageview"))}var f,v=window.history;function d(e){return e&&e.tagName&&"a"===e.tagName.toLowerCase()}v.pushState&&(f=v.pushState,v.pushState=function(){f.apply(this,arguments),l()},window.addEventListener("popstate",l)),"prerender"===a.visibilityState?a.addEventListener("visibilitychange",function(){s||"visible"!==a.visibilityState||l()}):l();var m=1;function g(e){if("auxclick"!==e.type||e.button===m){var t,n=function(e){for(;e&&(void 0===e.tagName||!d(e)||!e.href);)e=e.parentNode;return e}(e.target),r=n&&n.href&&n.href.split("?")[0];if(!function e(t,n){if(!t||x<n)return!1;if(S(t))return!0;return e(t.parentNode,n+1)}(n,0))return(t=n)&&t.href&&t.host&&t.host!==i.host?w(e,n,{name:"Outbound Link: Click",props:{url:n.href}}):function(e){if(!e)return!1;var t=e.split(".").pop();return k.some(function(e){return e===t})}(r)?w(e,n,{name:"File Download",props:{url:r}}):void 0}}function w(e,t,n){var r=!1;function i(){r||(r=!0,window.location=t.href)}!function(e,t){if(!e.defaultPrevented){var n=!t.target||t.target.match(/^_(self|parent|top)$/i),r=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type;return n&&r}}(e,t)?vince(n.name,{props:n.props}):(vince(n.name,{props:n.props,callback:i}),setTimeout(i,5e3),e.preventDefault())}a.addEventListener("click",g),a.addEventListener("auxclick",g);var h=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma"],y=o.getAttribute("file-types"),b=o.getAttribute("add-file-types"),k=y&&y.split(",")||b&&b.split(",").concat(h)||h;function L(e){var t=S(e)?e:e&&e.parentNode,n={name:null,props:{}},r=t&&t.classList;if(!r)return n;for(var i=0;i<r.length;i++){var a,o,p=r.item(i).match(/vince-event-(.+)=(.+)/);p&&(a=p[1],o=p[2].replace(/\+/g," "),"name"===a.toLowerCase()?n.name=o:n.props[a]=o)}return n}var x=3;function N(e){if("auxclick"!==e.type||e.button===m){for(var t,n,r,i,a=e.target,o=0;o<=x&&a;o++){if((r=a)&&r.tagName&&"form"===r.tagName.toLowerCase())return;d(a)&&(t=a),S(a)&&(n=a),a=a.parentNode}n&&(i=L(n),t?(i.props.url=t.href,w(e,t,i)):vince(i.name,{props:i.props}))}}function S(e){var t=e&&e.classList;if(t)for(var n=0;n<t.length;n++)if(t.item(n).match(/vince-event-name=(.+)/))return 1}a.addEventListener("submit",function(e){var t,n=e.target,r=L(n);function i(){t||(t=!0,n.submit())}r.name&&(e.preventDefault(),t=!1,setTimeout(i,5e3),vince(r.name,{props:r.props,callback:i}))}),a.addEventListener("click",N),a.addEventListener("auxclick",N)}();