import{f as c}from"./fetchData-47df6b6c.js";const d=async({url:s})=>{try{const e=s.searchParams.get("session_id");if(!e)throw new Error("Missing session ID");const r=await c({sessionId:e},"/orders/success");if(r.status!==200)throw new Error("There was an error retrieving Stripe session");const{customer:{name:t,address:o},customer_details:{email:a},payment_intent:{id:n},line_items:{data:[i]}}=await r.body;return{name:t,address:o,email:a,paymentId:n,product:i}}catch(e){console.error(`Error fetching checkout session: ${e}`)}},m=Object.freeze(Object.defineProperty({__proto__:null,load:d},Symbol.toStringTag,{value:"Module"}));export{m as _,d as l};