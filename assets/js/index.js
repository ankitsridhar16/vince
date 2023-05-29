"use strict";(()=>{var I=Object.defineProperty;var Z=Object.getOwnPropertyDescriptor;var s=(t,e)=>I(t,"name",{value:e,configurable:!0});var d=(t,e,o,r)=>{for(var n=r>1?void 0:r?Z(e,o):e,i=t.length-1,a;i>=0;i--)(a=t[i])&&(n=(r?a(e,o,n):a(n))||n);return r&&n&&I(e,o,n),n};var J=(t,e,o)=>{if(!e.has(t))throw TypeError("Cannot "+o)};var y=(t,e,o)=>{if(e.has(t))throw TypeError("Cannot add the same private member more than once");e instanceof WeakSet?e.add(t):e.set(t,o)};var D=(t,e,o)=>(J(t,e,"access private method"),o);var A=new WeakSet;function T(t){A.add(t),t.shadowRoot&&C(t.shadowRoot),S(t),k(t.ownerDocument)}s(T,"bind");function C(t){S(t),k(t)}s(C,"bindShadow");var v=new WeakMap;function k(t=document){if(v.has(t))return v.get(t);let e=!1,o=new MutationObserver(n=>{for(let i of n)if(i.type==="attributes"&&i.target instanceof Element)M(i.target);else if(i.type==="childList"&&i.addedNodes.length)for(let a of i.addedNodes)a instanceof Element&&S(a)});o.observe(t,{childList:!0,subtree:!0,attributeFilter:["data-action"]});let r={get closed(){return e},unsubscribe(){e=!0,v.delete(t),o.disconnect()}};return v.set(t,r),r}s(k,"listenForBind");function S(t){for(let e of t.querySelectorAll("[data-action]"))M(e);t instanceof Element&&t.hasAttribute("data-action")&&M(t)}s(S,"bindElements");function Q(t){let e=t.currentTarget;for(let o of W(e))if(t.type===o.type){let r=e.closest(o.tag);A.has(r)&&typeof r[o.method]=="function"&&r[o.method](t);let n=e.getRootNode();if(n instanceof ShadowRoot&&A.has(n.host)&&n.host.matches(o.tag)){let i=n.host;typeof i[o.method]=="function"&&i[o.method](t)}}}s(Q,"handleEvent");function*W(t){for(let e of(t.getAttribute("data-action")||"").trim().split(/\s+/)){let o=e.lastIndexOf(":"),r=Math.max(0,e.lastIndexOf("#"))||e.length;yield{type:e.slice(0,o),tag:e.slice(o+1,r),method:e.slice(r+1)||"handleEvent"}}}s(W,"bindings");function M(t){for(let e of W(t))t.addEventListener(e.type,Q)}s(M,"bindActions");var _=s(t=>String(typeof t=="symbol"?t.description:t).replace(/([A-Z]($|[a-z]))/g,"-$1").replace(/--/g,"-").replace(/^-|-$/,"").toLowerCase(),"dasherize"),$=s((t,e="property")=>{let o=_(t);if(!o.includes("-"))throw new DOMException(`${e}: ${String(t)} is not a valid ${e} name`,"SyntaxError");return o},"mustDasherize");function H(t){let e=_(t.name).replace(/-element$/,"");try{window.customElements.define(e,t),window[t.name]=customElements.get(e)}catch(o){if(!(o instanceof DOMException&&o.name==="NotSupportedError"))throw o}return t}s(H,"register");function P(t,e){let o=t.tagName.toLowerCase();if(t.shadowRoot){for(let r of t.shadowRoot.querySelectorAll(`[data-target~="${o}.${e}"]`))if(!r.closest(o))return r}for(let r of t.querySelectorAll(`[data-target~="${o}.${e}"]`))if(r.closest(o)===t)return r}s(P,"findTarget");function R(t,e){let o=t.tagName.toLowerCase(),r=[];if(t.shadowRoot)for(let n of t.shadowRoot.querySelectorAll(`[data-targets~="${o}.${e}"]`))n.closest(o)||r.push(n);for(let n of t.querySelectorAll(`[data-targets~="${o}.${e}"]`))n.closest(o)===t&&r.push(n);return r}s(R,"findTargets");function z(t){for(let e of t.querySelectorAll("template[data-shadowroot]"))e.parentElement===t&&t.attachShadow({mode:e.getAttribute("data-shadowroot")==="closed"?"closed":"open"}).append(e.content.cloneNode(!0))}s(z,"autoShadowRoot");var O="attr";var B=new WeakSet;function E(t,e){if(B.has(t))return;B.add(t);let o=Object.getPrototypeOf(t),r=o?.constructor?.attrPrefix??"data-";e||(e=g(o,O));for(let n of e){let i=t[n],a=$(`${r}${n}`),u={configurable:!0,get(){return this.getAttribute(a)||""},set(p){this.setAttribute(a,p||"")}};typeof i=="number"?u={configurable:!0,get(){return Number(this.getAttribute(a)||0)},set(p){this.setAttribute(a,p)}}:typeof i=="boolean"&&(u={configurable:!0,get(){return this.hasAttribute(a)},set(p){this.toggleAttribute(a,p)}}),Object.defineProperty(t,n,u),n in t&&!t.hasAttribute(a)&&u.set.call(t,i)}}s(E,"initializeAttrs");function q(t){let e=t.observedAttributes||[],o=t.attrPrefix??"data-",r=s(n=>$(`${o}${n}`),"attrToAttributeName");Object.defineProperty(t,"observedAttributes",{configurable:!0,get(){return[...g(t.prototype,O)].map(r).concat(e)},set(n){e=n}})}s(q,"defineObservedAttributes");var L=Symbol.for("catalyst"),b=class{constructor(e){let o=this,r=e.prototype.connectedCallback;e.prototype.connectedCallback=function(){o.connectedCallback(this,r)};let n=e.prototype.disconnectedCallback;e.prototype.disconnectedCallback=function(){o.disconnectedCallback(this,n)};let i=e.prototype.attributeChangedCallback;e.prototype.attributeChangedCallback=function(u,p,K){o.attributeChangedCallback(this,u,p,K,i)};let a=e.observedAttributes||[];Object.defineProperty(e,"observedAttributes",{configurable:!0,get(){return o.observedAttributes(this,a)},set(u){a=u}}),q(e),H(e)}observedAttributes(e,o){return o}connectedCallback(e,o){e.toggleAttribute("data-catalyst",!0),customElements.upgrade(e),z(e),E(e),T(e),o?.call(e),e.shadowRoot&&C(e.shadowRoot)}disconnectedCallback(e,o){o?.call(e)}attributeChangedCallback(e,o,r,n,i){E(e),o!=="data-catalyst"&&i&&i.call(e,o,r,n)}};s(b,"CatalystDelegate");function g(t,e){if(!Object.prototype.hasOwnProperty.call(t,L)){let r=t[L],n=t[L]=new Map;if(r)for(let[i,a]of r)n.set(i,new Set(a))}let o=t[L];return o.has(e)||o.set(e,new Set),o.get(e)}s(g,"meta");function l(t,e){g(t,"target").add(e),Object.defineProperty(t,e,{configurable:!0,get(){return P(this,e)}})}s(l,"target");function w(t,e){g(t,"targets").add(e),Object.defineProperty(t,e,{configurable:!0,get(){return R(this,e)}})}s(w,"targets");function h(t){new b(t)}s(h,"controller");var At=new Promise(t=>{document.readyState!=="loading"?t():document.addEventListener("readystatechange",()=>t(),{once:!0})}),Mt=new Promise(t=>{let e=new AbortController;e.signal.addEventListener("abort",()=>t());let o={once:!0,passive:!0,signal:e.signal},r=s(()=>e.abort(),"handler");document.addEventListener("mousedown",r,o),document.addEventListener("touchstart",r,o),document.addEventListener("keydown",r,o),document.addEventListener("pointerdown",r,o)});var N,U,c=class extends HTMLElement{constructor(){super(...arguments);y(this,N)}connectedCallback(){}changePeriod(o){var a;let r=o.target;this.periods.forEach(u=>{var p;(p=u.querySelector(".select-menu-item-icon"))==null||p.classList.add("d-none")});let n=r.parentElement;(a=n==null?void 0:n.querySelector(".select-menu-item-icon"))==null||a.classList.remove("d-none");let i=n==null?void 0:n.dataset;this.period_label.innerText=i==null?void 0:i.name,this.period_range_from_label.innerText=i==null?void 0:i.fromLabel,this.period_range_to_label.innerText=i==null?void 0:i.toLabel}changeMetrics(o){let r=o.target;this.metrics.forEach(i=>{var a;(a=i.parentElement)==null||a.classList.remove("navigation-focus")});let n=r.parentElement;n==null||n.classList.add("navigation-focus")}changeProps(o){let r=o.target;this.props.forEach(n=>{n.classList.remove("activeProps"),n.classList.add("border-0")}),r.classList.add("activeProps"),r.classList.remove("border-0")}};s(c,"VinceStatsElement"),N=new WeakSet,U=s(function(o){let r=Math.floor(Math.log10(Math.abs(o)));if(r<=2)return o.toString();let n=Math.floor(r/3),i=Math.pow(10,r-n*3)*+(o/Math.pow(10,r)).toFixed(1);return Math.round(i*100)/100+" "+["","K","M","B","T"][n]},"#shortNumber"),d([w],c.prototype,"periods",2),d([l],c.prototype,"period_label",2),d([l],c.prototype,"period_range_from_label",2),d([l],c.prototype,"period_range_to_label",2),d([w],c.prototype,"metrics",2),d([w],c.prototype,"props",2),c=d([h],c);window.customElements.get("vince-stats")||(window.VinceStatsElement=c,window.customElements.define("vince-stats",c));var x,F,m=class extends HTMLElement{constructor(){super(...arguments);y(this,x)}send(o){o.preventDefault(),o.stopImmediatePropagation();let r=document.createElement("form"),n=document.createElement("input");r.method=this.link.dataset.method,r.action=this.link.dataset.to,r.style.display="hidden",r.appendChild(D(this,x,F).call(this,"_csrf",this.link.dataset.csrf)),document.body.appendChild(r),n.type="submit",r.appendChild(n),console.log(r),n.click()}};s(m,"SendFormElement"),x=new WeakSet,F=s(function(o,r){var n=document.createElement("input");return n.type="hidden",n.name=o,n.value=r,n},"#buildHiddenInput"),d([l],m.prototype,"link",2),m=d([h],m);window.customElements.get("send-form")||(window.SendFormElement=m,window.customElements.define("send-form",m));var f=class extends HTMLElement{selectEvent(o){this.event.classList.add("propertySelected"),this.path.classList.remove("propertySelected"),this.event_fields.classList.remove("d-none"),this.path_fields.classList.add("d-none")}selectPath(o){this.path.classList.add("propertySelected"),this.event.classList.remove("propertySelected"),this.event_fields.classList.add("d-none"),this.path_fields.classList.remove("d-none")}};s(f,"GoalSelectionElement"),d([l],f.prototype,"event",2),d([l],f.prototype,"event_fields",2),d([l],f.prototype,"path",2),d([l],f.prototype,"path_fields",2),f=d([h],f);window.customElements.get("goal-selection")||(window.GoalSelectionElement=f,window.customElements.define("goal-selection",f));})();
