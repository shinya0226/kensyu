const express = require("express");
const app = express();
app.use(express.json());

//ローカルサーバ立ち上げ
app.listen(3000,console.log("サーバ開始されました。"));

app.get("/",(req,res)=>{
    res.send("プログラミングチュートリアルへようこそ");
});

//お客様情報をサーバに置いておく
const customers=[
    {title:"田中",id:1},
    {title:"斉藤",id:2},
    {title:"松本",id:3},
    {title:"山本",id:4},
    {title:"小林",id:5},
];
//データを取得できるようにする（GETメソッド）
app.get("/api/customers",(req,res) => {
    res.send(customers);
});

app.get("/api/customers/:id",(req,res)=>{
    const customer = customers.find((c)=>c.id===parseInt(req.params.id));
    res.send(customer);
})


//データを送信（作成）する（POSTメソッド）
app.post("/api/customers",(req,res)=>{
    const customer={
        title:req.body.title,
        id:customers.length+1,
    };
    customers.push(customer);
    res.send(customers);
})

//データを更新する（PUTメソッド）
app.put("/api/customers/:id",(req,res)=>{
    const customer = customers.find((c)=>c.id===parseInt(req.params.id));
    customer.title=req.body.title;
    res.send(customer);
}
);

//データを削除する（DELETEメソッド）
app.delete("/api/customers/:id",(req,res)=>{
    const customer = customers.find((c)=>c.id===parseInt(req.params.id));
    //打ち込んだ配列の取得
    const index = customers.indexOf(customer);
    //削除
    customers.splice(index,1);
    res.send(customer);
});