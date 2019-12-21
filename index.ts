import 'reflect-metadata';

import Koa from 'koa';
import Router from 'koa-router';

import Logger from 'koa-logger';
import Json from 'koa-json';

const app = new Koa();
const router = new Router();

app.context.db = {};

router.get('/', async (ctx: Koa.Context, next: Koa.Next) => {
	ctx.body = { hello: 'world' };

	await next();
});

app.use(Json());
app.use(Logger());

app.use(router.routes()).use(router.allowedMethods());

app.listen(8000, () => console.log('$ serving on port 8000'));
