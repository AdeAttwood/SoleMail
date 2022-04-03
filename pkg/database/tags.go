package database

import "strings"

// Applies a tag query to a list of tags returning the result of the query A tag
// query is a list of string that will add or remove tags from a tag list an
// example of a tag query is
//
// +one +two -three
//
// The above query will add the tags 'one' and 'two' and remove 'three' all the
// other existing tags will be unchanged
func ApplyTags(tag_list []string, tag_string string) []string {
	var tag_items = strings.Split(tag_string, " ")

	new_tags := make(map[string]bool)
	for _, tag_item := range tag_items {
		trimmed := strings.TrimSpace(tag_item)
		opp := trimmed[0:1]
		tag := trimmed[1:]

		if opp == "+" {
			new_tags[tag] = true
		} else {
			new_tags[tag] = false
		}
	}

	for _, tag := range tag_list {
		if _, found := new_tags[tag]; !found {
			new_tags[tag] = true
		}
	}

	tags := []string{}
	for tag, add := range new_tags {
		if add {
			tags = append(tags, tag)
		}
	}

	return tags
}
