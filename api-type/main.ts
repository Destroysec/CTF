import express, { Application, Request, Response } from 'express'

const app: Application = express()
const port:number =3301;
app.get('/:number/:vouchers', (req: Request, res: Response) => {
  let options:object = {
    method: 'POST',
    headers: {'Content-Type': 'application/json'},
    body: `{"mobile":"${req.params.number}"}`
  };
  let link:string = `https://gift.truemoney.com/campaign/vouchers/${req.params.vouchers}/redeem`
    fetch(link,options)
    .then(res => res.json())
    .then(json =>res.send(json))
    .catch(err => res.send(err));
})
app.listen(port, () => {
  console.log(`⚡️[server]: Server is running at http://localhost:${port}`);
});