"use strict";(()=>{var W=Object.defineProperty;var Q=Object.getOwnPropertyDescriptor;var s=(t,e)=>W(t,"name",{value:e,configurable:!0});var d=(t,e,o,r)=>{for(var n=r>1?void 0:r?Q(e,o):e,i=t.length-1,a;i>=0;i--)(a=t[i])&&(n=(r?a(e,o,n):a(n))||n);return r&&n&&W(e,o,n),n};var U=(t,e,o)=>{if(!e.has(t))throw TypeError("Cannot "+o)};var v=(t,e,o)=>{if(e.has(t))throw TypeError("Cannot add the same private member more than once");e instanceof WeakSet?e.add(t):e.set(t,o)};var u=(t,e,o)=>(U(t,e,"access private method"),o);var _=new WeakSet;function k(t){_.add(t),t.shadowRoot&&S(t.shadowRoot),H(t),$(t.ownerDocument)}s(k,"bind");function S(t){H(t),$(t)}s(S,"bindShadow");var A=new WeakMap;function $(t=document){if(A.has(t))return A.get(t);let e=!1,o=new MutationObserver(n=>{for(let i of n)if(i.type==="attributes"&&i.target instanceof Element)T(i.target);else if(i.type==="childList"&&i.addedNodes.length)for(let a of i.addedNodes)a instanceof Element&&H(a)});o.observe(t,{childList:!0,subtree:!0,attributeFilter:["data-action"]});let r={get closed(){return e},unsubscribe(){e=!0,A.delete(t),o.disconnect()}};return A.set(t,r),r}s($,"listenForBind");function H(t){for(let e of t.querySelectorAll("[data-action]"))T(e);t instanceof Element&&t.hasAttribute("data-action")&&T(t)}s(H,"bindElements");function X(t){let e=t.currentTarget;for(let o of O(e))if(t.type===o.type){let r=e.closest(o.tag);_.has(r)&&typeof r[o.method]=="function"&&r[o.method](t);let n=e.getRootNode();if(n instanceof ShadowRoot&&_.has(n.host)&&n.host.matches(o.tag)){let i=n.host;typeof i[o.method]=="function"&&i[o.method](t)}}}s(X,"handleEvent");function*O(t){for(let e of(t.getAttribute("data-action")||"").trim().split(/\s+/)){let o=e.lastIndexOf(":"),r=Math.max(0,e.lastIndexOf("#"))||e.length;yield{type:e.slice(0,o),tag:e.slice(o+1,r),method:e.slice(r+1)||"handleEvent"}}}s(O,"bindings");function T(t){for(let e of O(t))t.addEventListener(e.type,X)}s(T,"bindActions");var P=s(t=>String(typeof t=="symbol"?t.description:t).replace(/([A-Z]($|[a-z]))/g,"-$1").replace(/--/g,"-").replace(/^-|-$/,"").toLowerCase(),"dasherize"),R=s((t,e="property")=>{let o=P(t);if(!o.includes("-"))throw new DOMException(`${e}: ${String(t)} is not a valid ${e} name`,"SyntaxError");return o},"mustDasherize");function z(t){let e=P(t.name).replace(/-element$/,"");try{window.customElements.define(e,t),window[t.name]=customElements.get(e)}catch(o){if(!(o instanceof DOMException&&o.name==="NotSupportedError"))throw o}return t}s(z,"register");function B(t,e){let o=t.tagName.toLowerCase();if(t.shadowRoot){for(let r of t.shadowRoot.querySelectorAll(`[data-target~="${o}.${e}"]`))if(!r.closest(o))return r}for(let r of t.querySelectorAll(`[data-target~="${o}.${e}"]`))if(r.closest(o)===t)return r}s(B,"findTarget");function N(t,e){let o=t.tagName.toLowerCase(),r=[];if(t.shadowRoot)for(let n of t.shadowRoot.querySelectorAll(`[data-targets~="${o}.${e}"]`))n.closest(o)||r.push(n);for(let n of t.querySelectorAll(`[data-targets~="${o}.${e}"]`))n.closest(o)===t&&r.push(n);return r}s(N,"findTargets");function q(t){for(let e of t.querySelectorAll("template[data-shadowroot]"))e.parentElement===t&&t.attachShadow({mode:e.getAttribute("data-shadowroot")==="closed"?"closed":"open"}).append(e.content.cloneNode(!0))}s(q,"autoShadowRoot");var K="attr";var F=new WeakSet;function M(t,e){if(F.has(t))return;F.add(t);let o=Object.getPrototypeOf(t),r=o?.constructor?.attrPrefix??"data-";e||(e=w(o,K));for(let n of e){let i=t[n],a=R(`${r}${n}`),f={configurable:!0,get(){return this.getAttribute(a)||""},set(b){this.setAttribute(a,b||"")}};typeof i=="number"?f={configurable:!0,get(){return Number(this.getAttribute(a)||0)},set(b){this.setAttribute(a,b)}}:typeof i=="boolean"&&(f={configurable:!0,get(){return this.hasAttribute(a)},set(b){this.toggleAttribute(a,b)}}),Object.defineProperty(t,n,f),n in t&&!t.hasAttribute(a)&&f.set.call(t,i)}}s(M,"initializeAttrs");function I(t){let e=t.observedAttributes||[],o=t.attrPrefix??"data-",r=s(n=>R(`${o}${n}`),"attrToAttributeName");Object.defineProperty(t,"observedAttributes",{configurable:!0,get(){return[...w(t.prototype,K)].map(r).concat(e)},set(n){e=n}})}s(I,"defineObservedAttributes");var L=Symbol.for("catalyst"),E=class{constructor(e){let o=this,r=e.prototype.connectedCallback;e.prototype.connectedCallback=function(){o.connectedCallback(this,r)};let n=e.prototype.disconnectedCallback;e.prototype.disconnectedCallback=function(){o.disconnectedCallback(this,n)};let i=e.prototype.attributeChangedCallback;e.prototype.attributeChangedCallback=function(f,b,J){o.attributeChangedCallback(this,f,b,J,i)};let a=e.observedAttributes||[];Object.defineProperty(e,"observedAttributes",{configurable:!0,get(){return o.observedAttributes(this,a)},set(f){a=f}}),I(e),z(e)}observedAttributes(e,o){return o}connectedCallback(e,o){e.toggleAttribute("data-catalyst",!0),customElements.upgrade(e),q(e),M(e),k(e),o?.call(e),e.shadowRoot&&S(e.shadowRoot)}disconnectedCallback(e,o){o?.call(e)}attributeChangedCallback(e,o,r,n,i){M(e),o!=="data-catalyst"&&i&&i.call(e,o,r,n)}};s(E,"CatalystDelegate");function w(t,e){if(!Object.prototype.hasOwnProperty.call(t,L)){let r=t[L],n=t[L]=new Map;if(r)for(let[i,a]of r)n.set(i,new Set(a))}let o=t[L];return o.has(e)||o.set(e,new Set),o.get(e)}s(w,"meta");function p(t,e){w(t,"target").add(e),Object.defineProperty(t,e,{configurable:!0,get(){return B(this,e)}})}s(p,"target");function g(t,e){w(t,"targets").add(e),Object.defineProperty(t,e,{configurable:!0,get(){return N(this,e)}})}s(g,"targets");function y(t){new E(t)}s(y,"controller");var Ct=new Promise(t=>{document.readyState!=="loading"?t():document.addEventListener("readystatechange",()=>t(),{once:!0})}),_t=new Promise(t=>{let e=new AbortController;e.signal.addEventListener("abort",()=>t());let o={once:!0,passive:!0,signal:e.signal},r=s(()=>e.abort(),"handler");document.addEventListener("mousedown",r,o),document.addEventListener("touchstart",r,o),document.addEventListener("keydown",r,o),document.addEventListener("pointerdown",r,o)});var m,x,D,Y,c=class extends HTMLElement{constructor(){super(...arguments);v(this,m);v(this,D);this.state={}}connectedCallback(){}selectAllBox(o){u(this,m,x).call(this,o,this.all_box)}selectPagesBox(o){u(this,m,x).call(this,o,this.pages_box)}referrerBox(o){u(this,m,x).call(this,o,this.referrer_box)}countriesBox(o){u(this,m,x).call(this,o,this.countries_box)}osBox(o){u(this,m,x).call(this,o,this.os_box)}};s(c,"VinceStatsElement"),m=new WeakSet,x=s(function(o,r){let n=o.currentTarget;for(let i of r)i.classList.remove("propertySelected");n.classList.add("propertySelected")},"#select"),D=new WeakSet,Y=s(function(o){let r=Math.floor(Math.log10(Math.abs(o)));if(r<=2)return o.toString();let n=Math.floor(r/3),i=Math.pow(10,r-n*3)*+(o/Math.pow(10,r)).toFixed(1);return Math.round(i*100)/100+" "+["","K","M","B","T"][n]},"#shortNumber"),d([g],c.prototype,"all_box",2),d([g],c.prototype,"pages_box",2),d([g],c.prototype,"referrer_box",2),d([g],c.prototype,"countries_box",2),d([g],c.prototype,"os_box",2),c=d([y],c);window.customElements.get("vince-stats")||(window.VinceStatsElement=c,window.customElements.define("vince-stats",c));var C,Z,h=class extends HTMLElement{constructor(){super(...arguments);v(this,C)}send(o){o.preventDefault(),o.stopImmediatePropagation();let r=document.createElement("form"),n=document.createElement("input");r.method=this.link.dataset.method,r.action=this.link.dataset.to,r.style.display="hidden",r.appendChild(u(this,C,Z).call(this,"_csrf",this.link.dataset.csrf)),document.body.appendChild(r),n.type="submit",r.appendChild(n),console.log(r),n.click()}};s(h,"SendFormElement"),C=new WeakSet,Z=s(function(o,r){var n=document.createElement("input");return n.type="hidden",n.name=o,n.value=r,n},"#buildHiddenInput"),d([p],h.prototype,"link",2),h=d([y],h);window.customElements.get("send-form")||(window.SendFormElement=h,window.customElements.define("send-form",h));var l=class extends HTMLElement{selectEvent(o){this.event.classList.add("propertySelected"),this.path.classList.remove("propertySelected"),this.event_fields.classList.remove("d-none"),this.path_fields.classList.add("d-none")}selectPath(o){this.path.classList.add("propertySelected"),this.event.classList.remove("propertySelected"),this.event_fields.classList.add("d-none"),this.path_fields.classList.remove("d-none")}};s(l,"GoalSelectionElement"),d([p],l.prototype,"event",2),d([p],l.prototype,"event_fields",2),d([p],l.prototype,"path",2),d([p],l.prototype,"path_fields",2),l=d([y],l);window.customElements.get("goal-selection")||(window.GoalSelectionElement=l,window.customElements.define("goal-selection",l));})();
