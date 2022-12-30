### Hi there 👋

I'm Koen, I'm a passionate Software Engineer and DevOps / Automation Enthusiast. And I care about Developer Experience, "The best workflow is invisible".

Currently I'm very into Static Code Analysis, AST (treesitter), Mass refractoring using [comby](https://comby.dev) and Neovim Plugin creation

#### 🐱💻  Check out what I'm currently working on
{{range recentContributions 5}}
- [{{.Repo.Name}}]({{.Repo.URL}}) - {{.Repo.Description}} ({{humanize .OccurredAt}})
{{- end}}

#### 👀 My latest projects
{{range recentRepos 5}}
- [{{.Name}}]({{.URL}}) - {{.Description}}
{{- end}}

#### 📜 my recent blog posts
{{range rss "https://conradtheprogrammer.medium.com/feed" 5}}
- [{{.Title}}]({{.URL}}) ({{humanize .PublishedAt}})
{{- end}}

#### 📨 How to reach me

- Twitter: https://www.twitter.com/koenverburg
- Twitch: https://www.twitch.tv/conradtheprogrammer
- Instagram: https://www.instagram.com/conradtheprogrammer
- Site: https://conradtheprogrammer.com
- Medium: https://conradtheprogrammer.medium.com
