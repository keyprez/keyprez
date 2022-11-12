import{a6 as l}from"./index-f8865573.js";function $(t){return t<.5?4*t*t*t:.5*Math.pow(2*t-2,3)+1}function m(t){const o=t-1;return o*o*o+1}function g(t,{delay:o=0,duration:s=400,easing:c=l}={}){const a=+getComputedStyle(t).opacity;return{delay:o,duration:s,easing:c,css:r=>`opacity: ${r*a}`}}function b(t,{delay:o=0,duration:s=400,easing:c=m,x:a=0,y:r=0,opacity:e=0}={}){const n=getComputedStyle(t),f=+n.opacity,p=n.transform==="none"?"":n.transform,u=f*(1-e);return{delay:o,duration:s,easing:c,css:(y,i)=>`
			transform: ${p} translate(${(1-y)*a}px, ${(1-y)*r}px);
			opacity: ${f-u*i}`}}function x(t,{delay:o=0,duration:s=400,easing:c=m,start:a=0,opacity:r=0}={}){const e=getComputedStyle(t),n=+e.opacity,f=e.transform==="none"?"":e.transform,p=1-a,u=n*(1-r);return{delay:o,duration:s,easing:c,css:(y,i)=>`
			transform: ${f} scale(${1-p*i});
			opacity: ${n-u*i}
		`}}export{$ as a,g as b,m as c,b as f,x as s};
