# cloudmailin [![CircleCI](https://circleci.com/gh/orcaman/cloudmailin.svg?style=svg)](https://circleci.com/gh/orcaman/cloudmailin)
`cloudmailin` is an unofficial parser libary for `JSON` responses coming from [CloudMailin](https://www.cloudmailin.com/) service. 

## usage

See tests for more. Simple usage example (an email with txt attachment):

```go
package main

import (
	"github.com/orcaman/cloudmailin"

	"encoding/json"
	"log"
)

const cloudMailinResponse = `{
	"headers": {
		"Received": [
			"by mail-wm0-fgoogle.com with SMTP id r141so1514832wmg.1 for <13412321122ad019e9e57f07c03d8b@cloudmailin.net>; Thu, 02 Feb 2017 15:15:35 -0800",
			"by 10.80.142.30 with HTTP; Thu, 02 Feb 2017 15:15:14 -0800"
		],
		"Date": "Fri, 03 Feb 2017 01:15:14 +0200",
		"From": "Dan Brorn <danbrorn13123@gmail.com>",
		"To": "13412321122ad019e9e57f07c03d8b@cloudmailin.net",
		"Message-ID": "<CAH8+3pG1rfOeW_ceJ=0PuXWmB9mJTu+scpi-V=oMXL_=xENUfg@mail.gmail.com>",
		"Subject": "test",
		"Mime-Version": "1.0",
		"Content-Type": "multipart/mixed; boundary=f403045d59dea9902f05479457b1",
		"DKIM-Signature": "v=1; a=rsa-sha256; c=relaxed/relaxed; d=gmail.com; s=20161025; h=mime-version:from:date:message-id:subject:to; bh=dbDOKfEL6BwxCXIMQwxsET2ZmX0vKCmKcH3rT9jVGws=; b=QJ2UVoRiTSawHfLl0OapTodejZuv3zeDa7jE+dpCUIOes63hVXnNN+SOP8rjdVFk0a +8q+kcPeq1D0wlF7Y80ZqaSuQfeL7acjtyXiNuan+TUHem/zxHxvB/x7h8JIzjthPfsiQaWyeVFZCS1RbfaUDNFDMhkNEfFVSBo0R ah8NkE32PEQn3K+giqWC7YlfsjyX5A2RLaojWpS66DcdKxXlvAjUHPH0ltscan3rCL9U tlI6rH+QAVsJeWAHbgzS2NRQvMkeKmI6czdTXbCjKigq8ULwqbzQNk6ZThigUJaWCyZA nfQA==",
		"X-Google-DKIM-Signature": "v=1; a=rsa-sha256; c=relaxed/relaxed; d=1e100.net; s=20161025; h=x-gm-message-state:mime-version:from:date:message-id:subject:to; bh=dbDOKfEL6BwxCXIMQwxsET2ZmX0vKCmKcH3rT9jVGws=; b=BDvFitszX7PS68RogTVuODKDWeJz5USTdq4qbMA1w4SJTLr/K2b8IwEeX/bf542ekM /jkDRI7mrad8xogi9l0p+Hw djowvQ5TMxvPtwp4t74a6gxN2t/5eXRODQVLRDWX2N2hC0pj2e7m3XO+yUSxXUCnHu2Z s4wt3esTJbeKeHBYnuvwMjcm+OXP7ZQKtrGb7SEh+m6ZhDeJqrevmctjNWudEKun5uV4 nQ4rELiarc33+1vQXtHcdQS0YzScP5151zkO33oJfylyrnXwBgg3ZZg/3yoqPx9idpA6 iOVA==",
		"X-Gm-Message-State": "+SOIoACqK8r3hFc2qVG/xIJLKXJp9ZFIVjZ4x4bbOXA==",
		"X-Received": "by 10.223.142.1 with SMTP id n1mr9544575w; Thu, 02 Feb 2017 15:15:34 -0800 (PST)"
	},
	"envelope": {
		"to": "13412321122ad019e9e57f07c03d8b@cloudmailin.net",
		"recipients": [
			"13412321122ad019e9e57f07c03d8b@cloudmailin.net"
		],
		"from": "danbrorn13123@gmail.com",
		"helo_domain": "mail-333-f53.google.com",
		"remote_ip": "10.226.17.110",
		"spf": {
			"result": "neutral",
			"domain": "gmail.com"
		}
	},
	"plain": "attac\n",
	"html": "<div dir=\"ltr\">attac</div>\n",
	"attachments": [{
		"content": "YQpiCmMKZAo=",
		"file_name": "file.txt",
		"content_type": "text/plain",
		"size": "8",
		"disposition": "attachment"
	}]
}`

func main() {
	expectedJSON := []byte(cloudMailinResponse)

	var r *cloudmailin.Response

	err := json.Unmarshal(expectedJSON, &r)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("subject: %s", *r.Headers.Subject)
	log.Printf("attachments:\n%s (%s)", r.Attachments[0].Content, *r.Attachments[0].Filename)

	// prints out:
	// 2017/02/05 11:45:45 subject: test
	// 2017/02/05 11:45:45 attachments:
	// a
	// b
	// c
	// d
	// (file.txt)
}
```