// posts.go

...

func (
	rs postsResource) Routes() chi.Router {
		r := chi.NewRouter()

		// GET /posts - Read a list of posts.
		r.Get("/", rs.List)
		// POST /posts - Create a new post.
		r.Post("/", rs.Create)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(PostCtx)
			// GET /posts/{id} -
			// Read a single post by :id.
			r.Get("/", rs.Get)
			// PUT /posts/{id} -
			// Update a single post by :id
			r.Put("/" rs.Update)
			// DELETE /posts/{id} -
			// Delete a single Post by :id.
			r.Delete("/", rs.Delete)
		})

		return r
	}