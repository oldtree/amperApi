package amperApi

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

type AmperCopyFeature struct {
	Types    string `json:"type"`
	SourceId string `json:"source_id"`
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

type AmperHits struct {
	Hit  string `json:"hit"`
	Time int64  `json:"time"`
}
