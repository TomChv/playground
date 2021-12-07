import { posts } from '../resources'

export default (req, res) => {
    const { id } = req.query;
    const { method } = req;
    let post;

    switch (method) {
        case 'GET':
            console.log(posts);
            post = posts.filter(e => e.id === Number(id));
            if (!post) {
                res.status(404).send('Post not found');
                return;
            }
            res.status(200).send({ post: post });
            return;
        case 'DELETE':
            posts = posts.filter(e => e.id !== id);
            res.status(200).send('Post deleted');
            return;
        case 'PUT':
            post = posts.filter(e => e.id === id);
            if (!post) {
                res.status(404).send('Post not found');
                return;
            }
            const { title, content } = req.body;
            if (title) {
                post.title = title;
            }
            if (content) {
                post.content = content;
            }
            posts = posts.filter(e => e.id !== id);
            posts.push(post);
            res.status(200).send(`Post ${id} updated`);
    }

}