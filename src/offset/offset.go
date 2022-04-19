package offset

const (
	GAMETIME         int = 0x30ECC8C //F3 0F 11 05 ? ? ? ? 8B 49
	OBJECTMANAGER    int = 0x24A4B90 //89 ? ? ? ? ? 57 C7 06 ? ? ? ? 66 C7 46 04 ? ?
	LOCALPLAYER      int = 0x30F5BBC //51 8B 0D ? ? ? ? 85 C9 74 26
	UNDERMOUSEOBJECT int = 0x30E95F0 //8B 0D ? ? ? ? C7 04 24 ? ? ? ? FF 74 24 58 - 89 0D ? ? ? ? C7 41 ? ? ? ? ? C7 41 ? ? ? ? ? C7 01 ? ? ? ?
	ZOOMCLASS        int = 0x310D610 //A3 ? ? ? ? 83 FA 10 72 32
	CHAT             int = 0x30F5C94 //8B 0D ? ? ? ? 8A D8 85
	VIEWPROJMATRICES int = 0x31203A0 //B9 ? ? ? ? 0F 10 80
	RENDERER         int = 0x31232A0 //8B 0D ? ? ? ? 50 E8 ? ? ? ? 5E 8B
	MINIMAPOBJECT    int = 0x30ECCCC //8B 1D ? ? ? ? 85 DB 74 19

	OBJINDEX             int = 0x20  //always the same
	OBJTEAM              int = 0x4C  //always the same
	OBJMISSILENAME       int = 0x6C  //always the same
	OBJNETWORKID         int = 0xCC  //always the same
	OBJPOS               int = 0x23C //11.18
	OBJMISSILESPELLCAST  int = 0x250 //always the same
	OBJVISIBILITY        int = 0x28C
	OBJSPAWNCOUNT        int = 0x2A0
	OBJSRCINDEX          int = 0x2AC //always the same
	OBJMANA              int = 0x2B4
	OBJMAXMANA           int = 0x2C4
	OBJRECALLNAME        int = 0xD98
	OBJRECALLSTATE       int = OBJRECALLNAME + 0x10
	OBJHEALTH            int = 0xDB4
	OBJMAXHEALTH         int = OBJHEALTH + 0x1
	OBJABILITYHASTE      int = 0x16A0
	OBJLETHALITY         int = 0x11C0
	OBJARMOR             int = 0x12AC
	OBJBONUSARMOR        int = OBJARMOR + 0x4
	OBJMAGICRES          int = OBJARMOR + 0x8
	OBJBONUSMAGICRES     int = 0x12B8
	OBJBASEATK           int = 0x1284
	OBJBONUSATK          int = 0x11FC
	OBJMOVESPEED         int = 0x12C4
	OBJSPELLBOOK         int = 0x27C0
	OBJTRANSFORMATION    int = 0x3040 //always the same
	OBJNAME              int = 0x2BAC
	OBJLVL               int = 0x33A0
	OBJSIZEMULTIPLIER    int = 0x129C
	OBJEXPIRY            int = 0x298 //always the same
	OBJCRIT              int = 0x12E0
	OBJCRITMULTI         int = 0x12D0
	OBJABILITYPOWER      int = 0x1750
	OBJATKSPEEDMULTI     int = 0x1280
	OBJATKRANGE          int = 0x12CC
	OBJTARGETABLE        int = 0xD1C
	OBJINVULNERABLE      int = 0x3EC
	OBJISMOVING          int = 0x32EF //always the same
	OBJDIRECTION         int = 0x1C10 //0x1BD8L
	OBJITEMLIST          int = 0x33E8 //always the same
	OBJEXPIERIENCE       int = 0x3384
	OBJMAGICPEN          int = 0x11DC
	OBJMAGICPENMULTI     int = 0x11E4
	OBJADDITIONALAPMULTI int = 0x1230
	OBJMANAREGEN         int = 0x1150
	OBJHEALTHREGEN       int = 0x12F8
	OBJPERCENTARMORPEN   int = 0x11E0
	OBJPERCENTMAGICPEN   int = 0x11E4
	OBJLIFESTEAL         int = 0x12A0
	OBJCRITCHANCE        int = 0x12E0

	OBJEXPERIENCE int = 0x337C
	OBJLEVEL      int = OBJEXPERIENCE + 0x10

	OBJAVAIABLESPELLPOINTS int = 0x33CC

	OBJSUMMONERSPELL_D int = 0x3858
	OBJSUMMONERSPELL_F int = 0x3864
	OBJKEYSTONE        int = 0x3878

	MAXZOOM int = 0x20 //always the same

	CHATISOPEN int = 0x754 //C7 86 ? ? ? ? ? ? ? ? E8 ? ? ? ? 83 C4 04 85 C0 75 30 F6 86 ? ? ? ? ? 75 1B 38 86 ? ? ? ?

	SPELLBOOKACTIVESPELLCAST int = 0x20  //always the same
	SPELLBOOKSPELLSLOTS      int = 0x488 //always the same

	OBJBUFFMANAGER           int = 0x2180 //11.18,0x21B8 the ones below are always the same //8D 83 ? ? ? ? 50 8D AB ? ? ? ? // 4 first characters are the offset
	BUFFMANAGERENTRIESARRAY  int = 0x10
	BUFFMANAGERENDARRAY      int = 0x14
	BUFFENTRYBUFF            int = 0x8
	BUFFTYPE                 int = 0x4
	BUFFENTRYBUFFSTARTTIME   int = 0xC
	BUFFENTRYBUFFENDTIME     int = 0x10
	BUFFENTRYBUFFCOUNT       int = 0x74
	BUFFENTRYBUFFCOUNTALT    int = 0x24
	BUFFENTRYBUFFCOUNTALT2   int = 0x20
	BUFFNAME                 int = 0x8
	BUFFENTRYBUFFNODESTART   int = 0x20
	BUFFENTRYBUFFNODECURRENT int = 0x24

	//always the same
	ITEMLISTITEM int = 0xC
	ITEMINFO     int = 0x20
	ITEMINFOID   int = 0x68

	//always the same
	CURRENTDASHSPEED int = 0x1D0
	ISDASHING        int = 0x398
	DASHPOS          int = 0x1FC
	ISMOVING         int = 0x198
	NAVBEGIN         int = 0x1BC
	NAVEND           int = 0x1C0 //0x32B0 12.2

	//never change
	RENDERERWIDTH  int = 0xC
	RENDERERHEIGHT int = 0x10

	//spellslots never change
	SPELLSLOTLEVEL               int = 0x20
	SPELLSLOTTIME                int = 0x28
	SPELLSLOTCHARGES             int = 0x58
	SPELLSLOTTIMECHARGE          int = 0x78
	SPELLSLOTDAMAGE              int = 0x94
	SPELLSLOTSPELLINFO           int = 0x144
	SPELLINFOSPELLDATA           int = 0x44
	SPELLDATASPELLNAME           int = 0x6C
	SPELLDATAMISSILENAME         int = 0x6C
	SPELLSLOTSMITETIMER          int = 0x64
	SPELLSLOTSMITECHARGES        int = 0x58
	SPELLSLOTITEMMAXSTACKCOUNT   int = 0x6
	SPELLSLOTITEMNEXTREFILLTIME  int = 0x6
	SPELLSLOTITEMSTACKSTATE      int = 0x7
	SPELLSLOTITEMACTIVESTATE     int = 0x7
	SPELLSLOTITEMCOOLDOWN        int = 0x7
	SPELLSLOTITEMTARGETINGCLIENT int = 0x13
	SPELLSLOTITEMNAME            int = 0x13

	//these never change
	OBJECTMAPCOUNT      int = 0x2C
	OBJECTMAPROOT       int = 0x28
	OBJECTMAPNODENETID  int = 0x10
	OBJECTMAPNODEOBJECT int = 0x14

	//these never change
	SPELLCASTSPELLINFO    int = 0x8
	SPELLCASTSTARTTIME    int = 0x544
	SPELLCASTSTARTTIMEALT int = 0x534
	SPELLCASTCASTTIME     int = 0x4C0
	SPELLCASTSTART        int = 0x80
	SPELLCASTEND          int = 0x8C
	SPELLCASTSRCIDX       int = 0x68
	SPELLCASTDESTIDX      int = 0xC0

	MINIMAPOBJECTHUD int = 0x110
	MINIMAPHUDPOS    int = 0x44
	MINIMAPHUDSIZE   int = 0x4C

	// not in use, no need of updating
	AIMANAGER                   int = 0x2C98
	AIMANAGERSTARTPATH          int = 0x1CC // Funciona
	AIMANAGERENDPATH            int = 0x1D8 // Funciona
	AIMANAGERCLICKRIGHTPOSITION int = 0x10
	AIMANAGERCASTPOSITION       int = 0x388
	AIMANAGERISMOVING           int = 0x1C0 // Funciona
	AIMANAGEROWNPOSITION        int = 0x3E4 // Funciona
	AIMANAGERMOVESPEED          int = 0x1BC // Funciona
	AIMANAGERSERVPOSITION       int = 0x2E4 // Funciona
	AIMANAGERISDASHING          int = 0x3C0 //controlar con Ghost Spell
	AIMANAGERCURRENTSEGMENT     int = 0x1C4
	AIMANAGERDASHSPEED          int = 0x1F8 //Ghost Spell . si no viene el speed is not dashing
	AIMANAGERVELOCITY           int = 0x2F0
	AIMANAGERPOINTERPATHSTART   int = 0x1E4 //Vector3 list
	AIMANAGERPOINTERPATHEND     int = 0x1E8 //Vector3 list

	// TestGamePing
	TESTGAMEPINGA int = 0x310B7FC //A1 ?? ?? ?? ?? 85 C0 74 07 C7 40 ?? ?? ?? ?? ?? C2
	TESTGAMEPINGB int = 0x3c
	TESTGAMEPINGC int = 0x28

	MISSILEMAP         int = 0x34F848C //12.1
	MISSILEMAPCOUNT    int = 0x78      //12.1
	MISSILEMAPROOT     int = 0x74      //12.1
	MISSILEMAPKEY      int = 0x10      //12.1
	MISSILEMAPVAL      int = 0x14      //12.1
	MISSILESPELLINFO   int = 0x278     //0x8L //12.1
	MISSILESRCIDX      int = 0x2DC     //0x6CL //12.1
	MISSILEDESTIDX     int = 0x330     //0xC0L //12.1
	MISSILEDESTIDXTEST int = 0x334     //0xC0L //12.1
	MISSILESTARTPOS    int = 0x2F4     //0x90L //12.1
	MISSILEENDPOS      int = 0x300     //0x9CL //12.1
	MISSILESPEED       int = 0x68      //12.1

	AIMissileClient int = 0x24A4BD4
	AIMinionClient  int = 0x24A4C24
	AITurretClient  int = 0x30EA410
	AIHeroClient    int = 0x1856650
	AIBuildingList  int = 0x1855F10

// #define oTurretList 0x30EA410	//8B 35 ? ? ? ? 8B 76 18
// #define oMinionList 0x24A4C24	//A1 ? ? ? ? 8B 50 04 8D 48 04 89 44 24 3C
// #define oHeroList 0x1856650	//8B 15 ? ? ? ? 0F 44 C1
// #define oInhibitorList 0x30F5D28	//A1 ?? ?? ?? ?? 53 55 56 8B 70 04 8B 40 08
// #define oBuildingList 0x1855F10	//8B 15 ?? ?? ?? ?? 8B 7A 04 8B 42 08 8D 0C 87 3B F9 73 38
// #define oAttackableUnitList 0x24A4BC8	//A1 ?? ?? ?? ?? 8B 50 04 8B 40 08
// #define oShopsList 0x30EA414	//8B 0D ?? ?? ?? ?? E8 ?? ?? ?? ?? 84 C0 75 11 8B 4F 04
// #define oMissileList 0x24A4BD4 	//8B 35 ?? ?? ?? ?? 8B 46 04 8B 6E 08 8D 0C A8 3B C1 73 15 66 0F 1F 44 00 00 39 38
)
