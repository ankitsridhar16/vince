!function(){"use strict";var a=window.location,i=window.document,o=i.currentScript,c=o.getAttribute("data-api")||new URL(o.src).origin+"/api/event";function e(e,t){try{if("true"===window.localStorage.vince_ignore)return void console.warn("Ignoring Event: localStorage flag")}catch(e){}var n={};n.n=e,n.u=t&&t.u?t.u:a.href,n.d=o.getAttribute("data-domain"),n.r=i.referrer||null,n.w=window.innerWidth,t&&t.meta&&(n.m=JSON.stringify(t.meta)),t&&t.props&&(n.p=t.props);var r=new XMLHttpRequest;r.open("POST",c,!0),r.setRequestHeader("Content-Type","text/plain"),r.send(JSON.stringify(n)),r.onreadystatechange=function(){4===r.readyState&&t&&t.callback&&t.callback()}}var t=window.vince&&window.vince.q||[];window.vince=e;for(var n=0;n<t.length;n++)e.apply(this,t[n]);function p(e){return e&&e.tagName&&"a"===e.tagName.toLowerCase()}var u=1;function r(e){if("auxclick"!==e.type||e.button===u){var t,n=function(e){for(;e&&(void 0===e.tagName||!p(e)||!e.href);)e=e.parentNode;return e}(e.target),r=n&&n.href&&n.href.split("?")[0];if(!function e(t,n){if(!t||g<n)return!1;if(h(t))return!0;return e(t.parentNode,n+1)}(n,0))return(t=n)&&t.href&&t.host&&t.host!==a.host?s(e,n,{name:"Outbound Link: Click",props:{url:n.href}}):function(e){if(!e)return!1;var t=e.split(".").pop();return d.some(function(e){return e===t})}(r)?s(e,n,{name:"File Download",props:{url:r}}):void 0}}function s(e,t,n){var r=!1;function a(){r||(r=!0,window.location=t.href)}!function(e,t){if(!e.defaultPrevented){var n=!t.target||t.target.match(/^_(self|parent|top)$/i),r=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type;return n&&r}}(e,t)?vince(n.name,{props:n.props}):(vince(n.name,{props:n.props,callback:a}),setTimeout(a,5e3),e.preventDefault())}i.addEventListener("click",r),i.addEventListener("auxclick",r);var f=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma"],l=o.getAttribute("file-types"),v=o.getAttribute("add-file-types"),d=l&&l.split(",")||v&&v.split(",").concat(f)||f;function m(e){var t=h(e)?e:e&&e.parentNode,n={name:null,props:{}},r=t&&t.classList;if(!r)return n;for(var a=0;a<r.length;a++){var i,o,c=r.item(a).match(/vince-event-(.+)=(.+)/);c&&(i=c[1],o=c[2].replace(/\+/g," "),"name"===i.toLowerCase()?n.name=o:n.props[i]=o)}return n}var g=3;function w(e){if("auxclick"!==e.type||e.button===u){for(var t,n,r,a,i=e.target,o=0;o<=g&&i;o++){if((r=i)&&r.tagName&&"form"===r.tagName.toLowerCase())return;p(i)&&(t=i),h(i)&&(n=i),i=i.parentNode}n&&(a=m(n),t?(a.props.url=t.href,s(e,t,a)):vince(a.name,{props:a.props}))}}function h(e){var t=e&&e.classList;if(t)for(var n=0;n<t.length;n++)if(t.item(n).match(/vince-event-name=(.+)/))return 1}i.addEventListener("submit",function(e){var t,n=e.target,r=m(n);function a(){t||(t=!0,n.submit())}r.name&&(e.preventDefault(),t=!1,setTimeout(a,5e3),vince(r.name,{props:r.props,callback:a}))}),i.addEventListener("click",w),i.addEventListener("auxclick",w)}();