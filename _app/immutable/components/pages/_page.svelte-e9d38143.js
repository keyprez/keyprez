import{S as R,i as Y,s as Q,k as _,a as j,q as N,l as d,m as k,c as P,h,r as I,n as c,V as L,b as C,F as p,u as O,B as b,ab as T,w as B,x as D,y as G,ac as W,f as v,J as X,t as $,z as H,R as Z,e as U,g as ee,d as te,Y as re}from"../../chunks/index-f8865573.js";import{S as ae}from"../../chunks/SvelteSeo-233e266b.js";import{b as le}from"../../chunks/index-1785cdb6.js";import"../../chunks/store-e69f09a2.js";import{L as se}from"../../chunks/Loader-3cb8697c.js";import{f as oe}from"../../chunks/fetchProducts-73e560c4.js";function ne(n){let t,l,e,s,u,a,r,o,m,q,w,y=n[2].toUpperCase()+"",S,V,g,x,z;return{c(){t=_("div"),l=_("a"),e=_("img"),u=j(),a=_("span"),r=_("img"),q=j(),w=_("p"),S=N(y),V=j(),g=_("h2"),x=N(n[1]),z=N(" NOK"),this.h()},l(f){t=d(f,"DIV",{class:!0});var i=k(t);l=d(i,"A",{class:!0,"data-sveltekit-prefetch":!0,href:!0});var E=k(l);e=d(E,"IMG",{class:!0,src:!0,alt:!0}),u=P(E),a=d(E,"SPAN",{class:!0});var K=k(a);r=d(K,"IMG",{class:!0,src:!0,alt:!0}),K.forEach(h),E.forEach(h),q=P(i),w=d(i,"P",{class:!0});var M=k(w);S=I(M,y),M.forEach(h),V=P(i),g=d(i,"H2",{class:!0});var A=k(g);x=I(A,n[1]),z=I(A," NOK"),A.forEach(h),i.forEach(h),this.h()},h(){c(e,"class","object-cover opacity-100 w-full rounded-lg shadow-xl ease-linear duration-300 group-hover:opacity-20"),L(e.src,s=`${n[3].toLowerCase()}.jpg`)||c(e,"src",s),c(e,"alt",n[3]),c(r,"class","max-w-[15rem] min-w-[13rem] max-h-[10rem]"),L(r.src,o=`/${n[3].toLowerCase()}.svg`)||c(r,"src",o),c(r,"alt",n[3]),c(a,"class","opacity-0 absolute ease-linear duration-300 group-hover:opacity-100"),c(l,"class","flex flex-col justify-center items-center group"),c(l,"data-sveltekit-prefetch",""),c(l,"href",m=`/product/${n[0]}`),c(w,"class","text-black dark:text-white"),c(g,"class","dark:text-teal-500"),c(t,"class","text-center w-full flex-[1_0_100%] md:flex-[1_0_40%] lg:flex-[1_0_20%]")},m(f,i){C(f,t,i),p(t,l),p(l,e),p(l,u),p(l,a),p(a,r),p(t,q),p(t,w),p(w,S),p(t,V),p(t,g),p(g,x),p(g,z)},p(f,[i]){i&8&&!L(e.src,s=`${f[3].toLowerCase()}.jpg`)&&c(e,"src",s),i&8&&c(e,"alt",f[3]),i&8&&!L(r.src,o=`/${f[3].toLowerCase()}.svg`)&&c(r,"src",o),i&8&&c(r,"alt",f[3]),i&1&&m!==(m=`/product/${f[0]}`)&&c(l,"href",m),i&4&&y!==(y=f[2].toUpperCase()+"")&&O(S,y),i&2&&O(x,f[1])},i:b,o:b,d(f){f&&h(t)}}}function ce(n,t,l){let e,s,u,a,{product:r}=t;return n.$$set=o=>{"product"in o&&l(4,r=o.product)},n.$$.update=()=>{n.$$.dirty&16&&l(3,{name:e,description:s,price:u,slug:a}=r,e,(l(2,s),l(4,r)),(l(1,u),l(4,r)),(l(0,a),l(4,r)))},[a,u,s,e,r]}class ie extends R{constructor(t){super(),Y(this,t,ce,ne,Q,{product:4})}}function F(n,t,l){const e=n.slice();return e[2]=t[l],e}function ue(n){let t,l;return{c(){t=_("h1"),l=N("Something went wrong :(")},l(e){t=d(e,"H1",{});var s=k(t);l=I(s,"Something went wrong :("),s.forEach(h)},m(e,s){C(e,t,s),p(t,l)},p:b,i:b,o:b,d(e){e&&h(t)}}}function fe(n){let t,l,e=n[1],s=[];for(let a=0;a<e.length;a+=1)s[a]=J(F(n,e,a));const u=a=>$(s[a],1,1,()=>{s[a]=null});return{c(){for(let a=0;a<s.length;a+=1)s[a].c();t=U()},l(a){for(let r=0;r<s.length;r+=1)s[r].l(a);t=U()},m(a,r){for(let o=0;o<s.length;o+=1)s[o].m(a,r);C(a,t,r),l=!0},p(a,r){if(r&1){e=a[1];let o;for(o=0;o<e.length;o+=1){const m=F(a,e,o);s[o]?(s[o].p(m,r),v(s[o],1)):(s[o]=J(m),s[o].c(),v(s[o],1),s[o].m(t.parentNode,t))}for(ee(),o=e.length;o<s.length;o+=1)u(o);te()}},i(a){if(!l){for(let r=0;r<e.length;r+=1)v(s[r]);l=!0}},o(a){s=s.filter(Boolean);for(let r=0;r<s.length;r+=1)$(s[r]);l=!1},d(a){re(s,a),a&&h(t)}}}function J(n){let t,l;return t=new ie({props:{product:n[2]}}),{c(){B(t.$$.fragment)},l(e){D(t.$$.fragment,e)},m(e,s){G(t,e,s),l=!0},p:b,i(e){l||(v(t.$$.fragment,e),l=!0)},o(e){$(t.$$.fragment,e),l=!1},d(e){H(t,e)}}}function pe(n){let t,l;return t=new se({}),{c(){B(t.$$.fragment)},l(e){D(t.$$.fragment,e)},m(e,s){G(t,e,s),l=!0},p:b,i(e){l||(v(t.$$.fragment,e),l=!0)},o(e){$(t.$$.fragment,e),l=!1},d(e){H(t,e)}}}function he(n){let t,l,e,s,u;t=new ae({props:{title:"keyprez",description:"New keyboard shop in progress..."}});let a={ctx:n,current:null,token:null,hasCatch:!0,pending:pe,then:fe,catch:ue,value:1,blocks:[,,,]};return T(n[0],a),{c(){B(t.$$.fragment),l=j(),e=_("div"),a.block.c(),this.h()},l(r){D(t.$$.fragment,r),l=P(r),e=d(r,"DIV",{class:!0});var o=k(e);a.block.l(o),o.forEach(h),this.h()},h(){c(e,"class","flex content-center justify-between flex-wrap gap-6")},m(r,o){G(t,r,o),C(r,l,o),C(r,e,o),a.block.m(e,a.anchor=null),a.mount=()=>e,a.anchor=null,u=!0},p(r,[o]){n=r,W(a,n,o)},i(r){u||(v(t.$$.fragment,r),v(a.block),s||X(()=>{s=Z(e,le,{duration:1e3}),s.start()}),u=!0)},o(r){$(t.$$.fragment,r);for(let o=0;o<3;o+=1){const m=a.blocks[o];$(m)}u=!1},d(r){H(t,r),r&&h(l),r&&h(e),a.block.d(),a.token=null,a=null}}}function me(n){return[oe()]}class we extends R{constructor(t){super(),Y(this,t,me,he,Q,{})}}export{we as default};