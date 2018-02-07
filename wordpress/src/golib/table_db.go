package golib

import "time"

type TBCommentMeta struct {
	MetaId    int64  `orm:"column(meta_id);pk;auto"`
	CommentId int64  `orm:"column(comment_id);default(0);index"`
	MetaKey   string `orm:"column(meta_key);size(512);default("");index"`
	MetaValue string `orm:"column(meta_value);type(text);"`
}
type TBComments struct {
	CommentId          int64     `orm:"column(comment_id);pk;auto"`
	CommentPostId      int64     `orm:"column(comment_post_id);default(0)";index`
	CommentAuthor      string    `orm:"column(comment_author);type(text);"`                                                    //`comment_author` tinytext NOT NULL,
	CommentAuthorEmail string    `orm:"column(comment_author_email);size(100);index"`                                          //`comment_author_email` varchar(100) NOT NULL DEFAULT '',
	CommentAuthorUrl   string    `orm:"column(comment_author_url);size(200);"`                                                 //`comment_author_url` varchar(200) NOT NULL DEFAULT '',
	CommentAuthorIP    string    `orm:"column(comment_author_ip);size(200);"`                                                  //`comment_author_IP` varchar(100) NOT NULL DEFAULT '',
	CommentDate        time.Time `orm:"column(comment_date);auto_now;type(datetime);default("0000-00-00 00:00:00")"`           //`comment_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
	CommentDateGmt     time.Time `orm:"column(comment_date_gmt);auto_now;type(datetime);default("0000-00-00 00:00:00");index"` //`comment_date_gmt` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
	CommentContent     string    `orm:"column(comment_content);type(text);"`                                                   //`comment_content` text NOT NULL,
	CommentKarma       int64     `orm:"column(comment_karma);default(0);"`                                                     //`comment_karma` int(11) NOT NULL DEFAULT '0',
	CommentApproved    string    `orm:"column(comment_approved);size(20);default("0");index"`                                  //`comment_approved` varchar(20) NOT NULL DEFAULT '1',
	CommentAgent       string    `orm:"column(comment_agent);size(255);default("");"`                                          //`comment_agent` varchar(255) NOT NULL DEFAULT '',
	CommentType        string    `orm:"column(comment_type);size(20);default("");"`                                            //`comment_type` varchar(20) NOT NULL DEFAULT '',
	CommentParent      int64     `orm:"column(comment_parent);default(0);index"`                                               //`comment_parent` bigint(20) unsigned NOT NULL DEFAULT '0',
	UserId             int64     `orm:"column(user_id);default(0);"`                                                           //`user_id` bigint(20) unsigned NOT NULL DEFAULT '0',
}
type TBLinks struct {
	LinkId          int64     `orm:"column(link_id);pk;auto"`                                                     //`link_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	LinkUrl         string    `orm:"column(link_url);size(255);default("");"`                                     //`link_url` varchar(255) NOT NULL DEFAULT '',
	LinkName        string    `orm:"column(link_name);size(255);default("");"`                                    //`link_name` varchar(255) NOT NULL DEFAULT '',
	LinkImage       string    `orm:"column(link_image);size(255);default("");"`                                   //`link_image` varchar(255) NOT NULL DEFAULT '',
	LinkTarget      string    `orm:"column(link_target);size(25);default("");"`                                   //`link_target` varchar(25) NOT NULL DEFAULT '',
	LinkDescription string    `orm:"column(link_description);size(25);default("");"`                              //`link_description` varchar(255) NOT NULL DEFAULT '',
	LinkVisible     string    `orm:"column(link_visible);size(20);default("");index"`                             //`link_visible` varchar(20) NOT NULL DEFAULT 'Y',
	LinkOwner       int64     `orm:"column(link_owner);default(1);"`                                              //`link_owner` bigint(20) unsigned NOT NULL DEFAULT '1',
	LinkRating      int64     `orm:"column(link_rating);default(0);"`                                             //`link_rating` int(11) NOT NULL DEFAULT '0',
	LinkUpdated     time.Time `orm:"column(link_updated);auto_now;type(datetime);default("0000-00-00 00:00:00")"` //`link_updated` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
	LinkRel         string    `orm:"column(link_rel);size(255);default("");"`                                     //`link_rel` varchar(255) NOT NULL DEFAULT '',
	LinkNotes       string    `orm:"column(link_notes);size(255);type(mediumtext);default("");"`                  //`link_notes` mediumtext NOT NULL,
	LinkRss         string    `orm:"column(link_rss);size(255);default("");"`                                     //`link_rss` varchar(255) NOT NULL DEFAULT '',
}
type TBOptions struct {
	OptionId    int64  `orm:"column(option_id);pk;auto"`                       // `option_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	OptionName  string `orm:"column(option_name);size(191);default("");index"` //`option_name` varchar(191) NOT NULL DEFAULT '',
	OptionValue string `orm:"column(option_value);type(text);"`                //`option_value` longtext NOT NULL,
	AutoLoad    string `orm:"column(autoload);size(20);default("yes");"`       //`autoload` varchar(20) NOT NULL DEFAULT 'yes',
}
type TBPostMeta struct {
	MetaId    int64  `orm:"column(meta_id);pk;auto"`                      //`meta_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	PostId    int64  `orm:"column(post_id);default(0);index"`             //`post_id` bigint(20) unsigned NOT NULL DEFAULT '0',
	MetaKey   string `orm:"column(meta_key);size(255);default("");index"` //`meta_key` varchar(255) DEFAULT NULL,
	MetaValue string `orm:"column(meta_value);type(text);"`               //`meta_value` longtext,
}
type TBPost struct {
	Id                  int64     `orm:"column(id);pk;auto"`                                                               //`ID` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	PostAuthor          int64     `orm:"column(post_author);default(0);index"`                                             //`post_author` bigint(20) unsigned NOT NULL DEFAULT '0',
	PostDate            time.Time `orm:"column(post_date);auto_now;type(datetime);default("0000-00-00 00:00:00");index"`   //`post_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
	PostDateGmt         time.Time `orm:"column(post_date_gmt);auto_now;type(datetime);default("0000-00-00 00:00:00")"`     //`post_date_gmt` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
	PostContent         string    `orm:"column(post_content);type(text);"`                                                 //`post_content` longtext NOT NULL,
	PostTitle           string    `orm:"column(post_title);size(512);"`                                                    //post_title` text NOT NULL,
	PostExcerpt         string    `orm:"column(post_excerpt);size(512);"`                                                  //`post_excerpt` text NOT NULL,
	PostStatus          string    `orm:"column(post_status);size(20);index"`                                               //`post_status` varchar(20) NOT NULL DEFAULT 'publish',
	CommentStatus       string    `orm:"column(comment_status);size(20);"`                                                 //`comment_status` varchar(20) NOT NULL DEFAULT 'open',
	PingStatus          string    `orm:"column(ping_status);size(20);"`                                                    //`ping_status` varchar(20) NOT NULL DEFAULT 'open',
	PostPassword        string    `orm:"column(post_password);size(20);"`                                                  //`post_password` varchar(20) NOT NULL DEFAULT '',
	PostName            string    `orm:"column(post_name);size(200);index"`                                                //`post_name` varchar(200) NOT NULL DEFAULT '',
	ToPing              string    `orm:"column(to_ping);type(text);"`                                                      //`to_ping` text NOT NULL,
	Pinged              string    `orm:"column(pinged);type(text);"`                                                       //`pinged` text NOT NULL,
	PostModified        time.Time `orm:"column(post_modified);auto_now;type(datetime);default("0000-00-00 00:00:00")"`     //`post_modified` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
	PostModifiedGmt     time.Time `orm:"column(post_modified_gmt);auto_now;type(datetime);default("0000-00-00 00:00:00")"` //`post_modified_gmt` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
	PostContentFiltered string    `orm:"column(post_content_filtered);type(text);"`                                        //`post_content_filtered` longtext NOT NULL,
	PostParent          int64     `orm:"column(post_parent);default(0);index"`                                             //`post_parent` bigint(20) unsigned NOT NULL DEFAULT '0',
	Guid                string    `orm:"column(guid);size(255);"`                                                          //`guid` varchar(255) NOT NULL DEFAULT '',
	MenuOrder           int64     `orm:"column(menu_order);default(0)"`                                                    //`menu_order` int(11) NOT NULL DEFAULT '0',
	PostType            string    `orm:"column(post_type);size(255);index"`                                                //`post_type` varchar(20) NOT NULL DEFAULT 'post',
	PostMimeType        string    `orm:"column(post_mime_type);size(100);"`                                                //`post_mime_type` varchar(100) NOT NULL DEFAULT '',
	CommentCount        int64     `orm:"column(comment_count);default(0)"`                                                 //`comment_count` bigint(20) NOT NULL DEFAULT '0',
}
type TBTermMeta struct {
	MetaId    int64  `orm:"column(meta_id);pk;auto"`          //`meta_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	TermId    int64  `orm:"column(term_id);default(0);index"` // `term_id` bigint(20) unsigned NOT NULL DEFAULT '0',
	MetaKey   string `orm:"column(meta_key);size(255);index"` //`meta_key` varchar(255) DEFAULT NULL,
	MetaValue string `orm:"column(meta_value);type(text);"`   // `meta_value` longtext,
}
type TBTermRelationShips struct {
	ObjectId       int64 `orm:"column(object_id);pk;auto"`                    //`object_id` bigint(20) unsigned NOT NULL DEFAULT '0',
	TermTaxonomyId int64 `orm:"column(term_taxonomy_id);pk;default(0);index"` //`term_taxonomy_id` bigint(20) unsigned NOT NULL DEFAULT '0',
	TermOrder      int64 `orm:"column(term_order);default(0)"`                //`term_order` int(11) NOT NULL DEFAULT '0',
}
type TBTermTaxonomy struct {
	TermTaxonomyId int64  `orm:"column(term_taxonomy_id);pk;auto"` /// `term_taxonomy_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	TermId         int64  `orm:"column(term_id);default(0);index"` //`term_id` bigint(20) unsigned NOT NULL DEFAULT '0',
	Taxonomy       string `orm:"column(taxonomy);size(32);index"`  //`taxonomy` varchar(32) NOT NULL DEFAULT '',
	Description    string `orm:"column(description);type(text);"`  //`description` longtext NOT NULL,
	Parent         int64  `orm:"column(parent);default(0)"`        //`parent` bigint(20) unsigned NOT NULL DEFAULT '0',
	Count          int64  `orm:"column(count);default(0)"`         //`count` bigint(20) NOT NULL DEFAULT '0',
}
type TBUserMeta struct {
	UmetaId   int64  `orm:"column(umeta_id);pk;auto"`         //`umeta_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	UserId    int64  `orm:"column(user_id);default(0);index"` //`user_id` bigint(20) unsigned NOT NULL DEFAULT '0',
	MetaKey   string `orm:"column(meta_key);size(255);index"` //`meta_key` varchar(255) DEFAULT NULL,
	MetaValue string `orm:"column(meta_value);type(text);"`   //`meta_value` longtext,
}
type TBUser struct {
	id                int64     `orm:"column(id);pk;auto"`                                                             //`ID` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	UserLogin         string    `orm:"column(user_login);size(60);index"`                                              //`user_login` varchar(60) NOT NULL DEFAULT '',
	UserPass          string    `orm:"column(user_pass);size(255);"`                                                   //`user_pass` varchar(255) NOT NULL DEFAULT '',
	UserNicename      string    `orm:"column(user_nicename);size(50);index"`                                           //`user_nicename` varchar(50) NOT NULL DEFAULT '',
	UserEmail         string    `orm:"column(user_email);size(100);"`                                                  //`user_email` varchar(100) NOT NULL DEFAULT '',
	UserUrl           string    `orm:"column(user_url);size(100);"`                                                    //`user_url` varchar(100) NOT NULL DEFAULT '',
	UserRegistered    time.Time `orm:"column(user_registered);auto_now;type(datetime);default("0000-00-00 00:00:00")"` //`user_registered` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
	UserActivationKey string    `orm:"column(user_activation_key);size(100);"`                                         //`user_activation_key` varchar(255) NOT NULL DEFAULT '',
	UserStatus        int64     `orm:"column(user_status);default(0)"`                                                 //`user_status` int(11) NOT NULL DEFAULT '0',
	DisplayName       string    `orm:"column(display_name);size(250);"`                                                //`display_name` varchar(250) NOT NULL DEFAULT '',
}

func (p *TBCommentMeta) TableName() string       { return "commentmeta" }
func (p *TBComments) TableName() string          { return "comments" }
func (p *TBLinks) TableName() string             { return "links" }
func (p *TBOptions) TableName() string           { return "options" }
func (p *TBPostMeta) TableName() string          { return "postmeta" }
func (p *TBPost) TableName() string              { return "post" }
func (p *TBTermMeta) TableName() string          { return "termteta" }
func (p *TBTermRelationShips) TableName() string { return "termtelationships" }
func (p *TBTermTaxonomy) TableName() string      { return "termtaxonomy" }
func (p *TBUserMeta) TableName() string          { return "usermeta" }
func (p *TBUser) TableName() string              { return "user" }

func getAllTable() (instTable []interface{}) {
	instTable = append(instTable, new(TBCommentMeta))
	instTable = append(instTable, new(TBComments))
	instTable = append(instTable, new(TBLinks))
	instTable = append(instTable, new(TBOptions))
	instTable = append(instTable, new(TBPostMeta))
	instTable = append(instTable, new(TBPost))
	instTable = append(instTable, new(TBTermMeta))
	instTable = append(instTable, new(TBTermRelationShips))
	instTable = append(instTable, new(TBTermTaxonomy))
	instTable = append(instTable, new(TBUserMeta))
	instTable = append(instTable, new(TBUser))
	return
}
