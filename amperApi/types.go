package amperApi

import "time"

const ProjectNamingNote = "All object and property names, including all user-defined names, must be all lower case with underscores separating each word (i.e., lowercase_snake_case)."

type Amper interface {
}

//Unsuccessful or informational responses return a special status message
//If present, details can be an array containing human-readable message lines,
//or a set of validation objects that failed.

type AmperInfoMessages struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Details string `json:"details"`
}

type ApmerSingleObjectResp struct {
	Id      int    `jsno:"id"`
	PropOne string `json:"prop_one"`
	PropTwo string `json:"prop_two"`
}

type ApmerSingleObjectListResp struct {
	MultipleObject *ApmerSingleObjectResp `json:"objects"`
}

const (
	//Events define musical regions or periods of silence within a timeline. There are two types of events:

	//region events instruct the system to compose music for that region based on the time
	//and descriptor values, as well as optional attributes such as tempo and transition.
	//NOTE: Each region requires a unique user-defined numerical identifier.
	//This identifier only has to be unique within the project.
	//Simple one- or two-digit integer identifiers are usually sufficient for most projects.
	AmperEventsTypeRegion = "region"

	//silence events instruct the system to insert silence into the timeline
	//for the duration of that event.
	//Events are sequential.
	//One event ends when another begins.
	//For that reason, each timeline must end with a silence event.
	AmperEventsTypeSilence = "silence"
)

/*
Events define musical regions or periods of silence within a timeline. There are two types of events:
region
silence
region events instruct the system to compose music for that region based on the time and descriptor values,
 as well as optional attributes such as tempo and transition.
NOTE: Each region requires a unique user-defined numerical identifier.
 This identifier only has to be unique within the project.
 Simple one- or two-digit integer identifiers are usually sufficient for most projects.
silence events instruct the system to insert silence into the timeline for the duration of that event.
Events are sequential. One event ends when another begins.
 For that reason, each timeline must end with a silence event.
*/
const (
	AmperEventTypeRegion  = "region"
	AmperEventTypeSilence = "silence"
)

type AmperRegion struct {
	Event      string              `json:"event"`      //Text, required. This must be region.
	Id         int64               `jsno:"id"`         //Number, required. A unique identifier.
	Time       int64               `json:"time"`       //Number, required. The time, in seconds, at which the region will begin.
	Descriptor string              `json:"descriptor"` //Text, required. The key name for an internally recognized description of what mood, style, and other high level information needed to compose music.
	Tempo      int                 `jsno:"tempo"`      //Number, optional. The tempo, in beats per minute, the region will use. If omitted, an appropriate tempo will be chosen automatically.
	EndType    int                 `jsno:"end_type"`   //Object, optional. The end_type of the region. If ommited, an ending of type ringout will be used 2 seconds before the end of the region.
	Copy       []*AmperCopyFeature `jsno:"copy"`       //Array, optional. A list of all copy objects.
}

const (
	AmperCopyFeatureTypeHarmony = "harmony"
	AmperCopyFeatureTyperRythm  = "rhythm"
)

/*
In order to maintain rhythmic or harmonic continuity between regions,
you can specify that the rhythm or harmony in a region be copied from another specified region.
	type: Text, required. Which parameter of music to copy.
	source_id: Number required. The id of the region to copy from.
Valid options for type are:
	harmony
	rhythm
*/
type AmperCopyFeature struct {
	Types    string `json:"type"`
	SourceId string `json:"source_id"`
}

/*
The end_type of a region specifies how the ending portion of a region will be handled by the Composer.
There are two options: ending and transition.
	time: Number, required. The time, in seconds, at which the end_type will take effect.
	end_type: Text, required. Specifies whether the end_type is an ending or transition.
	type: Text, required. The specific action to perform for the end_type.
	          Valid options depend on which end_type the object is.

Valid ending types:
	ringout: At the given time, all instruments will play a final note,
	which will last until the region ends.

Valid transition types:
	tempo_ramp: At the given time, start a tempo ramp
			(from the tempo of the current region to the tempo of the next region)
			 which ends at the start of the next region.
	cut: Do nothing. The next region begins abruptly at the end of the current region.

Note that transitions are only valid if the current region is followed by an other region.
Invalid transitions will be converted to ringouts.
*/
type ApmerEndType struct {
	Time    int64  `json:"time"`
	EndType string `json:"end_type"`
	Type    string `json:"type"`
}

//AmperHitType

const (
	AmperHitType_chamber_crash      = "chamber_crash"
	AmperHitType_cymbal_crash       = "cymbal_crash"
	AmperHitType_cymbal_swell_1s    = "cymbal_swell_1s"
	AmperHitType_cymbal_swell_2s    = "cymbal_swell_2s"
	AmperHitType_cymbal_swell_500ms = "cymbal_swell_500ms"
	AmperHitType_deep_dry_hit       = "deep_dry_hit"
	AmperHitType_double_hammer      = "double_hammer"
	AmperHitType_dry_crunch_hit     = "dry_crunch_hit"
	AmperHitType_dry_flabby_hit     = "dry_flabby_hit"
	AmperHitType_dry_metal_hit_1    = "dry_metal_hit_1"
	AmperHitType_dry_metal_hit_2    = "dry_metal_hit_2"
	AmperHitType_flap_hit           = "flap_hit"
	AmperHitType_jingle_hit         = "jingle_hit"
	AmperHitType_punch_distortion   = "punch_distortion"
	AmperHitType_punch_tom          = "punch_tom"
	AmperHitType_suck_to_hit        = "suck_to_hit"
	AmperHitType_swirl_rise_and_hit = "swirl_rise_and_hit"
	AmperHitType_swish_to_hit       = "swish_to_hit"
	AmperHitType_tock_hit           = "tock_hit"
)

//Hits are optional statements, which allow for special musical emphasis at a specific point in time.
//		hit: Text, required. The name of the hit recognized by the Composer.
//		time: Number, required. The time at which the hit will occur.
type AmperHits struct {
	Hit  string `json:"hit"`
	Time int64  `json:"time"`
}

/*
Timelines define the musical regions for each project. Each timeline consists of:
	Event statements, which define musical regions or periods of silence
	Descriptors, which define the mood, style, etc. for each musical region.
	Optional tempo designations for each musical region. If a tempo is not explicitly defined, the system will use an internally generated tempo.
	Optional end_type instructions, which control how to transition from one musical region to the next, or from a musical region to a period of silence. If an end_type is not explicitly defined, the system will use an internally generated end_type.
	Optional hits which allow for special musical emphasis at a specific point in time.

Important considerations:
	The final event must be a silence.
	hit events can happen anywhere, including between silence events.
	The exact times of events may be altered slightly by the Composer to make them align with musical time.
	If a timeline is unparsable for any reason, the Composer will return a timeline of its default duration with a random descriptor.

We suggest limiting timelines to five minutes in length or less. Lengths greater than five minutes may not be supported.
*/

type AmperTimelines struct {
	TimelineSlice []*AmperRegion `json:"timelines"`
}

/*
Projects are a container for system interactions. They primarily define the musical timeline.

Projects can also accept an optional user-defined name.
If a user-defined name is not supplied, the system will generate one. However,
user-defined project names must be completely unique within that user account.
Duplicate project names will return an error.
*/
type AmperProject struct {
	Title                string          `jsno:"title"`
	AmperProjectTimeline *AmperTimelines `json:"timeline"`
}

type AmperDescriptor struct {
	Id   string `json:"id"`
	Name string `jsno:"name"`
}

type CreateAmperProjectResp struct {
	Id              int         `jsno:"id"`
	Status          string      `json:"status"`
	ProgressPercent string      `jsno:"progress_percent"`
	Files           interface{} `json:"files"`
	Created_at      time.Time   `json:"date_created"`
	Updated_at      time.Time   `jsno:"date_updated"`
}

type GetProjectInfoFile struct {
	Id          int       `jsno:"id"`
	ContentTye  string    `json:"content_type"`
	BitSample   int       `json:"bits_sample"`
	FrequencyHz int       `json:"frequency_hz"`
	KbitsSecond int       `json:"kbits_second"`
	DownloadUrl string    `json:"download_url"`
	SizeBytes   int       `json:"size_bytes"`
	Created_at  time.Time `json:"date_created"`
}
