import{S as C,i as v,s as w,k as h,l as p,m as y,h as f,n as m,b as x,G as B,E,g as S,t as d,d as q,f as k,q as L,r as N,F as T,u as j,B as b,w as z,x as A,y as F,z as G}from"./index-f8865573.js";import"./store-e69f09a2.js";import{L as O}from"./Loader-3cb8697c.js";function P(a){let e,n;return{c(){e=h("span"),n=L(a[0])},l(t){e=p(t,"SPAN",{});var l=y(e);n=N(l,a[0]),l.forEach(f)},m(t,l){x(t,e,l),T(e,n)},p(t,l){l&1&&j(n,t[0])},i:b,o:b,d(t){t&&f(e)}}}function U(a){let e,n;return e=new O({}),{c(){z(e.$$.fragment)},l(t){A(e.$$.fragment,t)},m(t,l){F(e,t,l),n=!0},p:b,i(t){n||(k(e.$$.fragment,t),n=!0)},o(t){d(e.$$.fragment,t),n=!1},d(t){G(e,t)}}}function D(a){let e,n,t,l,c,u;const o=[U,P],s=[];function g(i,r){return i[2]?0:1}return n=g(a),t=s[n]=o[n](a),{c(){e=h("button"),t.c(),this.h()},l(i){e=p(i,"BUTTON",{class:!0,type:!0});var r=y(e);t.l(r),r.forEach(f),this.h()},h(){m(e,"class","relative w-full md:max-w-xs flex justify-center bg-black dark:bg-teal-800 rounded-lg py-6 px-20 hover:bg-teal-900 dark:hover:bg-teal-900"),m(e,"type",a[1])},m(i,r){x(i,e,r),s[n].m(e,null),l=!0,c||(u=B(e,"click",function(){E(a[3])&&a[3].apply(this,arguments)}),c=!0)},p(i,[r]){a=i;let _=n;n=g(a),n===_?s[n].p(a,r):(S(),d(s[_],1,1,()=>{s[_]=null}),q(),t=s[n],t?t.p(a,r):(t=s[n]=o[n](a),t.c()),k(t,1),t.m(e,null)),(!l||r&2)&&m(e,"type",a[1])},i(i){l||(k(t),l=!0)},o(i){d(t),l=!1},d(i){i&&f(e),s[n].d(),c=!1,u()}}}function H(a,e,n){let{text:t}=e,{type:l}=e,{loading:c=!1}=e,{onClick:u}=e;return a.$$set=o=>{"text"in o&&n(0,t=o.text),"type"in o&&n(1,l=o.type),"loading"in o&&n(2,c=o.loading),"onClick"in o&&n(3,u=o.onClick)},[t,l,c,u]}class M extends C{constructor(e){super(),v(this,e,H,D,w,{text:0,type:1,loading:2,onClick:3})}}export{M as B};
