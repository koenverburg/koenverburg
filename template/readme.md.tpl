### Hi there 👋

I'm Koen, I'm a passionate Software Engineer and DevOps / automation enthusiast. I write small programs that make me more productive and I create boilerplates that focus on Developer Experience.

#### 🐱💻  Check out what I'm currently working on
{{range recentContributions 5}}
- [{{.Repo.Name}}]({{.Repo.URL}}) - {{.Repo.Description}} ({{humanize .OccurredAt}})
{{- end}}

#### 👀 My latest projects
{{range recentRepos 5}}
- [{{.Name}}]({{.URL}}) - {{.Description}}
{{- end}}

#### 📜 My recent blog posts
{{range rss "https://koenverburg.medium.com/feed" 5}}
- [{{.Title}}]({{.URL}}) ({{humanize .PublishedAt}})
{{- end}}

#### 📨 How to reach me

- Twitter: https://twitter.com/koenverburg_
- Twitch: https://www.twitch.tv/koenverburg
- Instagram (Tech): https://www.instagram.com/koenverburg.tech
- Site: https://koenverburg.dev
- Medium: https://koenverburg.medium.com
