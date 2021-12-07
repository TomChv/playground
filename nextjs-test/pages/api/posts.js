import { posts } from './resources'
let id = 0;


export default (req, res) => {
    const { method } = req;

    switch (method) {
        case 'GET':
           res.status(200).send({ posts });
           break;
        case 'POST':
            const { title, content } = req.body;
            if (!title || !content) {
                res.status(400).send('Invalid body.');
                return;
            }
            posts.push( { id, title, content, created_at: new Date()});
            id += 1;
            res.status(201).send('Posts successfully created.');
            break;
        default:
            res.setHeader('Allow', ['GET', 'PUT'])
            res.status(405).end(`Method ${method} not allowed`)
    }
}