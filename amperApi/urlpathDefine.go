package amperApi

const (
	/*
		Return a list of all Descriptors currently supported.
		The response object contains a descriptors property, which is an array of zero or more Descriptor objects.
		The following response fields are present:
			id: Text. The machine-readable identifier of this Descriptor. This is what must be used when constructing Timelines.
			name: Text. A human-readable string explaining this Descriptor. This is appropriate for user-facing UI elements.
	*/
	GetAmperDescriptors = "/v1/data/descriptors"

	/*
		Create a new Project in the authenticated user's account and initiate a Render.
		The request payload contains a Timeline and an optional title, both of which will be incorporated into the Render.
		The title, if present, must be unique across the user's entire account. If a title is not provided, a unique one will be generated.
		The following request fields are supported:
			timeline: Object, required. A structure describing musical qualities that are desired. The format and available options are described elsewhere in this document.
			title: Text. The title of this Project. If provided, must be unique across the user's entire account. If not provided, a unique title will be generated.
	*/
	CreateAmperProject = "/v1/projects"

	/*
		Query a Project for its Render status and location of files (if they are ready).
		The response object contains information about a Project's Render status.
		Additionally, if the Render completed successfully,
		the response will contain information about what Render files were created and where they can be downloaded.
	*/
	GetProjectInfo = "/v1/projects/:project_id"

	/*
		Download a single Render audio file. Note that this endpoint may return an HTTP redirect to a different domain, and the client must follow any and all redirects encountered to get the final file.
		The exact response varies depending on the content.
		Generally, after redirects are followed,
		the response will consist of binary data adhering to the content_type in the file object where this URL was listed.
	*/
	CetRenderFileInfo = "/v1/render_files/:render_file_id"
)

type GetAmperDescriptorsOperation struct {
	Descriptors []*AmperDescriptor `json:"descriptors"`
}

type CreateAmperProjectOperation struct {
	Project *AmperProject `jsno:"project"`
}
