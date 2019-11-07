package util

import (
	"os"
	"testing"
)

const (
	text1 = `新华社河内11月12日电（记者孙浩　孟娜　侯丽军）中共中央总书记、国家主席习近平12日在河内越共中央驻地同越共中央总书记阮富仲举行会谈。
	双方积极评价彼此发展成就，一致同意继承、维护、发扬好中越传统友谊，相互坚定奉行友好政策，把两国全面战略合作发展得越来越好，
	给两国和两国人民带来更多切实利益。习近平和阮富仲分别介绍了中共十九大、越南党和国家事业有关情况。习近平强调，中越是山水相连的好邻居、
	荣辱与共的好朋友、志同道合的好同志、合作共赢的好伙伴。中越建交67年来，由毛泽东主席、胡志明主席等两国老一辈领导人亲手缔造和精心培育的中
	越传统友谊不断得到巩固、弘扬。今年以来，两党两国高层保持密切接触，各领域合作不断深化，中越关系取得积极进展。当前，中越两国均处在改革发展
	的关键阶段，两国各自建设事业和双边关系发展面临难得的历史性机遇。双方要以两党两国关系大局为重，从两国人民根本利益出发，弘扬风雨同舟的优
	良传统，不断开创中越关系新局面，在前进道路上和衷共济、共谋发展。

	习近平对越南实施革新开放30多年来各项建设取得的巨大成就表示祝贺。习近平强调，相信在以阮富仲总书记为首的越共中央坚强领导下，兄弟的越南人
	民一定能够开创社会主义建设事业更加美好的未来。阮富仲衷心祝贺中共十九大取得圆满成功，祝贺习近平同志再次当选中共中央总书记、并被确立为中
	国共产党的领导核心。阮富仲表示，越方感谢中方对越南民族解放和国家发展给予的巨大帮助，支持中国在国际和地区事务中发挥更大作用。越方高度评
	价中国经济社会发展成就和对本地区发展的积极贡献，欢迎习近平同志提出的构建人类命运共同体主张，认为这体现了中国的全球视野和大国担当。越方
	愿同中方一道，深化两国传统友谊和全面战略合作伙伴关系，造福两国人民，促进地区和平与繁荣。

	习近平和阮富仲就新形势下深化中越全面战略合作伙伴关系达成重要共识，一致认为中越两国是有着悠久友好传统的邻国，都是共产党领导的社会主义
	国家，政治制度相同、发展道路相近，前途相关、命运与共。双方要相互借鉴，共同发展，为各自国家社会主义建设事业注入新活力，推动中越全面战略
	合作伙伴关系持续健康稳定向前发展，为促进地区和平、稳定、繁荣作出积极贡献。

	双方同意保持和加强高层交往优良传统，相互坚定奉行友好政策，加强战略沟通，增进政治互信，妥善处理分歧，引领中越关系正确方向。深化党际交往，
	发挥中越关系独特优势，密切治党治国理政经验交流。充分发挥中越双边合作指导委员会作用，加强外交、国防、公安、安全等各部门各层级交流合作。

	双方同意落实好共建“一带一路”和“两廊一圈”合作文件，促进地区经济联系和互联互通，推动经贸、产能、投资、基础设施建设、货币金融等领域合作
	不断取得务实进展，稳步推进跨境经济合作区建设，加强农业、环境、科技、交通运输等领域合作。

	双方同意扩大文化、教育、媒体、卫生、青年、地方、旅游等领域合作，推动河内中国文化中心、越中友谊宫尽早投入使用，办好民间交往活动，夯实中
	越友好民意基础。

	双方同意按照两党两国领导人达成的重要共识，妥善处理海上问题，稳步推进包括共同开发在内的各种形式的海上合作，共同致力于维护南海的和平与稳定。

	双方同意加强在联合国、世界贸易组织、亚太经合组织、亚欧会议、中国－东盟、澜沧江－湄公河合作等国际和地区框架内的配合。中方祝贺越方成功
	主办亚太经合组织第二十五次领导人非正式会议，愿同包括越南在内的各方一道，为推动亚太区域经济一体化、促进地区发展繁荣作出积极贡献。

	会谈后，两国领导人共同见证了共建“一带一路”和“两廊一圈”合作备忘录以及产能、能源、跨境经济合作区、电子商务、人力资源、经贸、金融、文化、
	卫生、新闻、社会科学、边防等领域合作文件的签署。

	会谈前，阮富仲在主席府广场举行盛大欢迎仪式。广场外簇拥着身着民族服装迎候习近平的当地民众。习近平乘车抵达时，越南少年儿童挥舞着中越两国
	国旗热烈欢迎习近平到来。阮富仲在停车处迎接，越南儿童向习近平献上鲜花。习近平同阮富仲登上检阅台。军乐队奏中越两国国歌，鸣21响礼炮。习近
	平在阮富仲陪同下检阅仪仗队。全体仪仗队员用越南语高呼“祝总书记身体健康”。习近平同越方陪同人员握手，阮富仲同中方陪同人员握手。习近平和阮
	富仲返回检阅台，观看分列式。两国领导人前往越共中央驻地时，越南青年抛撒花瓣欢迎。

	丁薛祥、刘鹤、杨洁篪等出席上述活动。越南方面出席的有：越共中央政治局委员、中央书记处常务书记、中央检查委员会主任陈国旺，越共中央政治局
	委员、中央书记处书记、中央组织部部长范明政，越共中央政治局委员、副总理、外交部部长范平明，越共中央政治局委员、中央书记处书记、中央宣教部
	部长武文赏，越共中央政治局委员、国防部部长吴春历，越共中央政治局委员、公安部部长苏林，越共中央书记处书记、中央办公厅主任阮文年及越南政
	府有关部门负责人。`

	text2 = `[
    {
        "data": {
            "Ontology name": "MPI Golm"
        },
        "parent": "#",
        "text": "MPI Golm",
        "id": "MPI Golm"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "测试一下这是什么玩意",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Asparagine",
        "id": "MPI Golm:ASPARAGINE"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "Averaged value of the Aspartate content",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Aspartate",
        "id": "MPI Golm:ASPARTATE"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "Averaged value of the Brix degree",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Brix",
        "id": "MPI Golm:BRIX"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "Averaged value of the Erythritol content",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Erythritol",
        "id": "MPI Golm:ERYTHRITOL"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "Averaged value of the Fucose content",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Fucose",
        "id": "MPI Golm:FUCOSE"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the GABA content",
            "Short name": "GABA",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "GABA",
        "id": "MPI Golm:GABA"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "Averaged value of the Malate content",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Malate",
        "id": "MPI Golm:MALATE"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "Averaged value of the Nicotinate content",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Nicotinate",
        "id": "MPI Golm:NICOTINATE"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "Averaged value of the Phenylalanine content",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Phenylalanine",
        "id": "MPI Golm:PHENYLALANINE"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "Averaged value of the Proline content",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Proline",
        "id": "MPI Golm:PROLINE"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "Averaged value of the Rhamnose content",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Rhamnose",
        "id": "MPI Golm:RHAMNOSE"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "Averaged value of the Threonate content",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Threonate",
        "id": "MPI Golm:THREONATE"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "Averaged value of the Threonine content",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Threonine",
        "id": "MPI Golm:THREONINE"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "Averaged value of the Tocopherol content",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Tocopherol",
        "id": "MPI Golm:TOCOPHEROL"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "Averaged value of the Citrate content",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Citrate",
        "id": "MPI Golm:CITRATE"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": null,
            "Description": "Averaged value of the Tyramine content",
            "Short name": null,
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Tyramine",
        "id": "MPI Golm:TYRAMINE"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Asparagine content",
            "Short name": "Asparagine",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Asparagine",
        "id": "MPI Golm:Asparagine"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Aspartate content",
            "Short name": "Aspartate",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Aspartate",
        "id": "MPI Golm:Aspartate"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the beta Alanine content",
            "Short name": "betaAlanine",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "betaAlanine",
        "id": "MPI Golm:betaAlanine"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Brix degree",
            "Short name": "Brix",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Brix",
        "id": "MPI Golm:Brix"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Citrate content",
            "Short name": "Citrate",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Citrate",
        "id": "MPI Golm:Citrate"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Stevens, R., Buret, M., Garchery, C., Carretero, Y. & Causse, M. Technique for Rapid, Small-Scale Analysis of Vitamin C Levels in Fruit and Application to a Tomato Mutant Collection. J. Agric. Food Chem. 54, 6159¿6165 (2006).",
            "Description": "Averaged value of the Dehydro-ascorbate content",
            "Short name": "DHA",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "DHA",
        "id": "MPI Golm:DHA"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Erythritol content",
            "Short name": "Erythritol",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Erythritol",
        "id": "MPI Golm:Erythritol"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Fucose content",
            "Short name": "Fucose",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Fucose",
        "id": "MPI Golm:Fucose"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Galacturonate content",
            "Short name": "Galacturonate",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Galacturonate",
        "id": "MPI Golm:Galacturonate"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Glucuronate content",
            "Short name": "Glucuronate",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Glucuronate",
        "id": "MPI Golm:Glucuronate"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Glutamate content",
            "Short name": "Glutamate",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Glutamate",
        "id": "MPI Golm:Glutamate"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Glutamine content",
            "Short name": "Glutamine",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Glutamine",
        "id": "MPI Golm:Glutamine"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Glutarate2oxo content",
            "Short name": "Glutarate2oxo",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Glutarate2oxo",
        "id": "MPI Golm:Glutarate2oxo"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Glycerol3P content",
            "Short name": "Glycerol3P",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Glycerol3P",
        "id": "MPI Golm:Glycerol3P"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Inositol1P content",
            "Short name": "Inositol1P",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Inositol1P",
        "id": "MPI Golm:Inositol1P"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Lysine content",
            "Short name": "Lysine",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Lysine",
        "id": "MPI Golm:Lysine"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Malate content",
            "Short name": "Malate",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Malate",
        "id": "MPI Golm:Malate"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Maltitol content",
            "Short name": "Maltitol",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Maltitol",
        "id": "MPI Golm:Maltitol"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Maltose content",
            "Short name": "Maltose",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Maltose",
        "id": "MPI Golm:Maltose"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Methionine content",
            "Short name": "Methionine",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Methionine",
        "id": "MPI Golm:Methionine"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Nicotinate content",
            "Short name": "Nicotinate",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Nicotinate",
        "id": "MPI Golm:Nicotinate"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Phenylalanine content",
            "Short name": "Phenylalanine",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Phenylalanine",
        "id": "MPI Golm:Phenylalanine"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Proline content",
            "Short name": "Proline",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Proline",
        "id": "MPI Golm:Proline"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Putrescine content",
            "Short name": "Putrescine",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Putrescine",
        "id": "MPI Golm:Putrescine"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Rhamnose content",
            "Short name": "Rhamnose",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Rhamnose",
        "id": "MPI Golm:Rhamnose"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Saccharate content",
            "Short name": "Saccharate",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Saccharate",
        "id": "MPI Golm:Saccharate"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Serine content",
            "Short name": "Serine",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Serine",
        "id": "MPI Golm:Serine"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Threonate content",
            "Short name": "Threonate",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Threonate",
        "id": "MPI Golm:Threonate"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Threonine content",
            "Short name": "Threonine",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Threonine",
        "id": "MPI Golm:Threonine"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Tocopherol content",
            "Short name": "Tocopherol",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Tocopherol",
        "id": "MPI Golm:Tocopherol"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Tyramine content",
            "Short name": "Tyramine",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Tyramine",
        "id": "MPI Golm:Tyramine"
    },
    {
        "data": {
            "Unit": null,
            "Measurement protocol": "Lisec, J., Schauer, N., Kopka, J., Willmitzer, L. & Fernie, A.R. Gas chromatography mass spectrometry¿based metabolite profiling in plants. Nat Protoc 1, 387¿396 (2006).",
            "Description": "Averaged value of the Xylose content",
            "Short name": "Xylose",
            "Ontology name": "MPI Golm"
        },
        "parent": "MPI Golm",
        "text": "Xylose",
        "id": "MPI Golm:Xylose"
    }
]`
	text3 = `{"head":{"v":10000,"rinfo":"{\"bizType\":0,\"brandID\":4,\"closeRoomTime\":0,\"delaySeconds\":300,\"enableRDP\":0,\"endTime\":1511949701,\"fakeOnlineNumber\":0,\"isRecord\":0,\"muteChat\":0,\"originalKMSLocation\":\"SH-TMS-ROOM\",\"originalKMSServerIP\":\"172.16.4.189:8888\",\"originalKMSServerId\":108,\"rdpAccount\":\"\",\"realKMSLocation\":\"SH-TMS-ROOM\",\"realKMSServerIP\":\"172.16.4.189:8888\",\"realKMSServerId\":0,\"roomNum\":\"00000286\",\"roomStatus\":1,\"roomType\":3,\"sessionId\":\"2017112917808186\",\"sessionRoomId\":\"2017112917808186_000286_ZTf3SIAL_ucRGwy\",\"startTime\":1511949101,\"videoInfo\":{\"bitrate\":100,\"frameRate\":20,\"height\":480,\"width\":640}}","minfo":["{\"identityKey\":\"student_nEBCdvdFfS1\",\"nickName\":\"student_nEBCdvdFfS1\",\"role\":10,\"token\":\"696d3f67450a40598542b112023925ec\"}","{\"identityKey\":\"student_nEBCdvdFfS2\",\"nickName\":\"student_nEBCdvdFfS2\",\"role\":10,\"token\":\"a4aa3f09691848d5bf736f0fa3f5ef33\"}","{\"identityKey\":\"student_nEBCdvdFfS3\",\"nickName\":\"student_nEBCdvdFfS3\",\"role\":10,\"token\":\"4aa7022cd431471cbcf55b68c149e6e6\"}","{\"identityKey\":\"teacher_nEBCdvdFfS\",\"nickName\":\"teacher_nEBCdvdFfS\",\"role\":20,\"token\":\"81e1e9f703174518bab24c4eb31152e6\"}"]},"events":[{"ts":1511949105,"sid":7,"ssid":1,"body":"{\"token\":\"81e1e9f703174518bab24c4eb31152e6\",\"name\":\"teacher_nEBCdvdFfS\",\"role\":20,\"audio\":true,\"video\":true,\"os\":\"1:10.13.1:1:62.0.3202.94:2\",\"region\":\"CN\"}"},{"ts":1511949106,"sid":7,"ssid":1,"body":"{\"token\":\"696d3f67450a40598542b112023925ec\",\"name\":\"student_nEBCdvdFfS1\",\"role\":10,\"audio\":true,\"os\":\"1:10.13.1:1:62.0.3202.94:2\",\"region\":\"CN\"}"},{"ts":1511949106,"sid":7,"ssid":2,"body":"{\"token\":\"81e1e9f703174518bab24c4eb31152e6\"}"},{"ts":1511949106,"sid":7,"ssid":10,"body":"{\"key\":\"pageId\",\"value\":\"14e972b2d94e30f2571e2f033c0ca7f5.png\"}"},{"ts":1511949107,"sid":7,"ssid":90,"body":"[{\"key\":\"switch\",\"value\":1}]"},{"ts":1511949107,"sid":7,"ssid":10,"body":"{\"key\":\"pageId\",\"value\":\"14e972b2d94e30f2571e2f033c0ca7f5.png\"}"},{"ts":1511949107,"sid":7,"ssid":2,"body":"{\"token\":\"81e1e9f703174518bab24c4eb31152e6\"}"},{"ts":1511949108,"sid":7,"ssid":2,"body":"{\"token\":\"696d3f67450a40598542b112023925ec\"}"},{"ts":1511949108,"sid":7,"ssid":1,"body":"{\"token\":\"81e1e9f703174518bab24c4eb31152e6\",\"name\":\"teacher_nEBCdvdFfS\",\"role\":20,\"audio\":true,\"video\":true,\"os\":\"1:10.13.1:1:62.0.3202.94:2\",\"region\":\"CN\"}"},{"ts":1511949109,"sid":7,"ssid":1,"body":"{\"token\":\"696d3f67450a40598542b112023925ec\",\"name\":\"student_nEBCdvdFfS1\",\"role\":10,\"audio\":true,\"os\":\"1:10.13.1:1:62.0.3202.94:2\",\"region\":\"CN\"}"},{"ts":1511949114,"sid":7,"ssid":11,"body":"[{\"key\":\"token\",\"value\":\"696d3f67450a40598542b112023925ec\"},{\"key\":\"camera\",\"value\":1}]"},{"ts":1511949117,"sid":4,"ssid":1,"body":"[[8,9117558090,[397.8571428571429,145],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949117,"sid":4,"ssid":40,"body":"[9117558090,\"s\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949117,"sid":4,"ssid":40,"body":"[9117558090,\"sd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949118,"sid":4,"ssid":40,"body":"[9117558090,\"sdf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949118,"sid":4,"ssid":40,"body":"[9117558090,\"sdfs\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949118,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949118,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949118,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949118,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949118,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\ns\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949118,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\nsf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949119,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\nsfj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949119,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\nsfjk\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949119,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\nsfjks\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949119,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\nsfjksd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949119,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\nsfjksdj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949119,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\nsfjksdjf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949119,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\nsfjksdjf\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949120,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\nsfjksdjf\\n黄金时代回房间\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949121,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\nsfjksdjf\\n黄金时代回房间\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949121,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\nsfjksdjf\\n黄金时代回房间\\n速度发货君\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949121,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\nsfjksdjf\\n黄金时代回房间\\n速度发货君\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949123,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\nsfjksdjf\\n黄金时代回房间\\n速度发货君\\nh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949124,"sid":4,"ssid":40,"body":"[9117558090,\"sdfsdfhk\\nsfjksdjf\\n黄金时代回房间\\n速度发货君\\nh九块\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949127,"sid":4,"ssid":1,"body":"[[8,9127469940,[249.2857142857143,292.8571428571429],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949128,"sid":4,"ssid":40,"body":"[9127469940,\"水淀粉你们呢\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949131,"sid":4,"ssid":1,"body":"[[8,9131158340,[415.7142857142857,279.28571428571433],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949131,"sid":4,"ssid":1,"body":"[[8,9131710810,[423.5714285714286,288.57142857142856],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949132,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949133,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949135,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949135,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949136,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949137,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949138,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949138,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949138,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949138,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949138,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjs\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949138,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949138,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd \",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949139,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949139,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949140,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949140,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949140,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjs\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949140,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjsh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949140,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949140,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949140,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949140,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj \",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949141,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj \\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949142,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj \\n黄金时代回房间\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949143,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj \\n黄金时代回房间\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949143,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj \\n黄金时代回房间\\n速度回房间\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949143,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj \\n黄金时代回房间\\n速度回房间\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949144,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj \\n黄金时代回房间\\n速度回房间\\n双方很快说\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949145,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj \\n黄金时代回房间\\n速度回房间\\n双方很快说\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949146,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj \\n黄金时代回房间\\n速度回房间\\n双方很快说\\n速度发货君\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949146,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj \\n黄金时代回房间\\n速度回房间\\n双方很快说\\n速度发货君\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949147,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj \\n黄金时代回房间\\n速度回房间\\n双方很快说\\n速度发货君\\n速度发货就好\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949147,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj \\n黄金时代回房间\\n速度回房间\\n双方很快说\\n速度发货君\\n速度发货就好\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949148,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj \\n黄金时代回房间\\n速度回房间\\n双方很快说\\n速度发货君\\n速度发货就好\\n速度回房间\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949148,"sid":4,"ssid":40,"body":"[9131710810,\"到沙发上的好风景速度恢复计划将黄金时代好\\n是的肌肤\\n还是觉得放寒假\\nhjsd f\\nhjshdfj \\n黄金时代回房间\\n速度回房间\\n双方很快说\\n速度发货君\\n速度发货就好\\n速度回房间\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949149,"sid":4,"ssid":1,"body":"[[8,9149406640,[481.42857142857144,309.2857142857143],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949151,"sid":4,"ssid":1,"body":"[[8,9151116060,[147.85714285714286,102.85714285714286],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949152,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949153,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949154,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949155,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949155,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949156,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949157,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949157,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949157,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949158,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949159,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949160,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949160,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949161,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949161,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949161,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949161,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949162,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949162,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949162,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949163,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949163,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949164,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949164,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949165,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\\n速度发货君\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949165,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\\n速度发货君dsf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949166,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\\n速度发货君dsf\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949167,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\\n速度发货君dsf\\n速度发货君\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949167,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\\n速度发货君dsf\\n速度发货君\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949167,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\\n速度发货君dsf\\n速度发货君\\n速度发货君\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949168,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\\n速度发货君dsf\\n速度发货君\\n速度发货君sd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949168,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\\n速度发货君dsf\\n速度发货君\\n速度发货君sd放寒假\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949168,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\\n速度发货君dsf\\n速度发货君\\n速度发货君sd放寒假\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949169,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\\n速度发货君dsf\\n速度发货君\\n速度发货君sd放寒假\\n都放寒假\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949169,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\\n速度发货君dsf\\n速度发货君\\n速度发货君sd放寒假\\n都放寒假dsf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949170,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\\n速度发货君dsf\\n速度发货君\\n速度发货君sd放寒假\\n都放寒假dsf烦得很近\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949170,"sid":4,"ssid":40,"body":"[9151116060,\"谁都放水淀粉好久好久十点回家\\n速度发货君就好\\n回家的时候房间\\n速度发货君\\n速度发货君速度发货君速度发货君\\n但是放寒假\\n水淀粉黄金时代分\\n发生点胡椒粉\\n烦得很；还是冬季发货\\n速度发货君\\n速度发货君dsf\\n速度发货君\\n速度发货君sd放寒假\\n都放寒假dsf烦得很近\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949185,"sid":4,"ssid":1,"body":"[[8,9185080000,[287.85714285714283,405.00000000000006],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949185,"sid":4,"ssid":1,"body":"[[8,9185767730,[155,333.5714285714286],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949187,"sid":4,"ssid":40,"body":"[9185767730,\"俄方的遗憾\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949188,"sid":4,"ssid":40,"body":"[9185767730,\"俄方的遗憾速度恢复计划将\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949188,"sid":4,"ssid":40,"body":"[9185767730,\"俄方的遗憾速度恢复计划将 \",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949189,"sid":4,"ssid":40,"body":"[9185767730,\"俄方的遗憾速度恢复计划将 \\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949189,"sid":4,"ssid":40,"body":"[9185767730,\"俄方的遗憾速度恢复计划将 \\n速度发货君回家舒服\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949190,"sid":4,"ssid":40,"body":"[9185767730,\"俄方的遗憾速度恢复计划将 \\n速度发货君回家舒服还是觉得放寒假\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949191,"sid":4,"ssid":40,"body":"[9185767730,\"俄方的遗憾速度恢复计划将 \\n速度发货君回家舒服还是觉得放寒假 \",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949191,"sid":4,"ssid":40,"body":"[9185767730,\"俄方的遗憾速度恢复计划将 \\n速度发货君回家舒服还是觉得放寒假 \\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949191,"sid":4,"ssid":40,"body":"[9185767730,\"俄方的遗憾速度恢复计划将 \\n速度发货君回家舒服还是觉得放寒假 \\n速度发货就好\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949191,"sid":4,"ssid":40,"body":"[9185767730,\"俄方的遗憾速度恢复计划将 \\n速度发货君回家舒服还是觉得放寒假 \\n速度发货就好\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949193,"sid":7,"ssid":10,"body":"{\"key\":\"pageId\",\"value\":\"4ebf5a7996ee0d52ca28ed259f13512a.png\"}"},{"ts":1511949194,"sid":4,"ssid":1,"body":"[[8,9194621800,[214.2857142857143,155],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949195,"sid":4,"ssid":40,"body":"[9194621800,\"水淀粉速度发货君\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949195,"sid":4,"ssid":40,"body":"[9194621800,\"水淀粉速度发货君\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949196,"sid":4,"ssid":40,"body":"[9194621800,\"水淀粉速度发货君\\n速度恢复计划将\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949199,"sid":4,"ssid":1,"body":"[[8,9199705920,[349.2857142857143,283.5714285714286],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949201,"sid":4,"ssid":40,"body":"[9194621800,\"水淀粉速度发货君\\n速度恢复计划将谁都放水淀粉\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949202,"sid":4,"ssid":40,"body":"[9185767730,\"俄方的遗憾速度恢复计划将 谁都放水淀粉\\n速度发货君回家舒服还是觉得放寒假 \\n速度发货就好\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949205,"sid":4,"ssid":1,"body":"[[8,9205409980,[405.7142857142857,360.7142857142858],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949206,"sid":4,"ssid":40,"body":"[9205409980,\"谁都放水淀粉\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949207,"sid":4,"ssid":1,"body":"[[8,9207884710,[454.28571428571433,262.85714285714283],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949209,"sid":4,"ssid":1,"body":"[[8,9209392590,[191.42857142857144,312.8571428571429],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949210,"sid":4,"ssid":40,"body":"[9209392590,\"谁都放水淀粉\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949217,"sid":4,"ssid":1,"body":"[[8,9217739390,[416.42857142857144,137.14285714285714],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949218,"sid":4,"ssid":1,"body":"[[8,9218043140,[416.42857142857144,137.14285714285714],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949219,"sid":4,"ssid":40,"body":"[9218043140,\"谁都放水淀粉\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949219,"sid":4,"ssid":1,"body":"[[8,9219539320,[468.5714285714286,183.57142857142858],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949235,"sid":4,"ssid":1,"body":"[[8,9235958910,[266.42857142857144,475],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949239,"sid":4,"ssid":40,"body":"[9209392590,\"谁都放水淀粉更好地睡个好觉\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949240,"sid":4,"ssid":40,"body":"[9209392590,\"谁都放水淀粉更好地睡个好觉速度给发货哈\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949241,"sid":4,"ssid":40,"body":"[9209392590,\"谁都放水淀粉更好地睡个好觉速度给发货哈水淀粉更好\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949242,"sid":4,"ssid":1,"body":"[[8,9242793640,[261.42857142857144,437.14285714285717],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949244,"sid":4,"ssid":40,"body":"[9242793640,\"谁都放水淀粉\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949248,"sid":7,"ssid":10,"body":"{\"key\":\"pageId\",\"value\":\"14e972b2d94e30f2571e2f033c0ca7f5.png\"}"},{"ts":1511949250,"sid":4,"ssid":1,"body":"[[8,9250242440,[241.42857142857144,340.7142857142857],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949251,"sid":4,"ssid":1,"body":"[[8,9251816730,[452.14285714285717,380],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949252,"sid":4,"ssid":1,"body":"[[8,9252334920,[482.8571428571429,425.00000000000006],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949255,"sid":4,"ssid":1,"body":"[[8,9255790210,[526.4285714285714,237.14285714285714],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949262,"sid":4,"ssid":1,"body":"[[8,9262773050,[477.14285714285717,420.7142857142857],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949264,"sid":4,"ssid":1,"body":"[[8,9264495300,[237.14285714285714,277.14285714285717],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949265,"sid":4,"ssid":1,"body":"[[8,9265982430,[435.7142857142858,185.71428571428575],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949266,"sid":4,"ssid":1,"body":"[[8,9266453040,[431.42857142857144,185.71428571428575],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949268,"sid":4,"ssid":40,"body":"[9266453040,\"水淀粉撒\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949270,"sid":4,"ssid":40,"body":"[9266453040,\"水淀粉撒谁的粉丝多\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949272,"sid":4,"ssid":1,"body":"[[8,9272910980,[445,365.00000000000006],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949273,"sid":4,"ssid":1,"body":"[[8,9273281600,[467.85714285714283,372.14285714285717],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949273,"sid":4,"ssid":1,"body":"[[8,9274000980,[234.2857142857143,325.00000000000006],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949274,"sid":4,"ssid":1,"body":"[[8,9274417120,[372.8571428571429,364.28571428571433],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949276,"sid":4,"ssid":1,"body":"[[8,9276273680,[475.7142857142858,417.8571428571429],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949277,"sid":7,"ssid":10,"body":"{\"key\":\"pageId\",\"value\":\"4ebf5a7996ee0d52ca28ed259f13512a.png\"}"},{"ts":1511949278,"sid":4,"ssid":1,"body":"[[8,9278714380,[392.14285714285717,305.7142857142857],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949279,"sid":4,"ssid":1,"body":"[[8,9279413060,[438.57142857142856,307.8571428571429],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949280,"sid":4,"ssid":1,"body":"[[8,9280379110,[464.2857142857143,350.7142857142857],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949281,"sid":4,"ssid":1,"body":"[[8,9281921620,[307.8571428571429,425.00000000000006],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949283,"sid":4,"ssid":40,"body":"[9281921620,\"水淀粉撒\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949285,"sid":4,"ssid":1,"body":"[[8,9285307480,[318.5714285714286,146.42857142857144],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949286,"sid":4,"ssid":40,"body":"[9194621800,\"水淀粉速度发货君是\\n速度恢复计划将谁都放水淀粉\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949287,"sid":4,"ssid":1,"body":"[[8,9288041050,[312.8571428571429,153.57142857142858],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949288,"sid":4,"ssid":40,"body":"[9288041050,\" \",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949289,"sid":4,"ssid":1,"body":"[[8,9289379090,[357.14285714285717,157.14285714285717],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949289,"sid":4,"ssid":1,"body":"[[8,9289808080,[372.8571428571429,158.57142857142856],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949290,"sid":4,"ssid":1,"body":"[[8,9290932510,[314.28571428571433,157.85714285714286],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949291,"sid":4,"ssid":1,"body":"[[8,9291383390,[266.42857142857144,152.14285714285717],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949292,"sid":4,"ssid":40,"body":"[9291383390,\"但是放速度\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949294,"sid":4,"ssid":1,"body":"[[8,9294298180,[222.85714285714286,165.71428571428572],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949325,"sid":4,"ssid":20,"body":"{\"ids\":[9209392590],\"offset\":{\"x\":-71.42857142857144,\"y\":-22.142857142857167}}"},{"ts":1511949328,"sid":7,"ssid":2,"body":"{\"token\":\"696d3f67450a40598542b112023925ec\"}"},{"ts":1511949329,"sid":7,"ssid":1,"body":"{\"token\":\"696d3f67450a40598542b112023925ec\",\"name\":\"student_nEBCdvdFfS1\",\"role\":10,\"audio\":true,\"video\":true,\"os\":\"1:10.13.1:1:62.0.3202.94:2\",\"region\":\"CN\"}"},{"ts":1511949332,"sid":4,"ssid":1,"body":"[[8,9332265160,[492.14285714285717,227.14285714285717],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949333,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949333,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949334,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949334,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949335,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\ns\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949335,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949335,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949335,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949335,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949335,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949335,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949335,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\ns\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949335,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949335,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949335,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949335,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949335,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjs\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949335,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\ns\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfs\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjs\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949336,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949337,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949337,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949337,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949337,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\ns\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949337,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949337,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949337,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949337,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhs\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949337,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949337,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949337,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949337,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949338,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949338,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949338,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\ns\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949338,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949338,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949338,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949338,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949338,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjs\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949338,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949338,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949338,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949339,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh \",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949339,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh \\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949341,"sid":4,"ssid":1,"body":"[[8,9341209750,[602.8571428571429,356.42857142857144],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949342,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh s\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949342,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sd\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949342,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdf\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949342,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfs\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949342,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsd\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949342,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdf\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949342,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfs\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949342,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949343,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\n\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949343,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\ns\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949343,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsd\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949343,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdf\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949343,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfs\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949343,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsd\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949343,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949343,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\n\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949344,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\ns\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949344,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsd\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949344,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdf\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949344,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949344,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\n\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949344,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\nd\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949344,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\nds\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949344,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949345,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\n\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949345,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\ns\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949345,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\nsd\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949345,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\nsdf\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949345,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\nsdf\\n\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949345,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\nsdf\\ns\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949346,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\nsdf\\nss\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949346,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\nsdf\\nssd\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949346,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\nsdf\\nssdf\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949346,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\nsdf\\nssdf\\n\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949346,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\nsdf\\nssdf\\ns\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949346,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\nsdf\\nssdf\\nsd\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949346,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\nsdf\\nssdf\\nsd\\n\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949346,"sid":4,"ssid":40,"body":"[9332265160,\"谁的粉丝多\\n水淀粉更黄金时代好\\nsdfhj \\nsdfhjsdf \\nsdfsdhjfjsdhfjk \\nsdfhsdhfj \\nsdfhjsdfh sdfsdfsd\\nsdfsdf\\nsdfs\\ndsf\\nsdf\\nssdf\\nsd\\nf\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949411,"sid":4,"ssid":1,"body":"[[8,9411432970,[601.4285714285714,100.71428571428572],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949411,"sid":4,"ssid":40,"body":"[9411432970,\"d\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949411,"sid":4,"ssid":40,"body":"[9411432970,\"ds\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949412,"sid":4,"ssid":40,"body":"[9411432970,\"dsf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949412,"sid":4,"ssid":40,"body":"[9411432970,\"dsfh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949412,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhg\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949412,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949412,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949413,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\ns\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949413,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949413,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949413,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949413,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949413,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949414,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949414,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\ns\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949414,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nsh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949414,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949414,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949414,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949414,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949415,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949415,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949415,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\ns\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949415,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949415,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949415,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949415,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949416,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949416,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949416,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\ns\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949416,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949416,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949416,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949416,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949417,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949417,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949417,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\ns\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949417,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\nsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949417,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\nsdf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949417,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\nsdfj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949418,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\nsdfjh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949418,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\nsdfjhs\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949418,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\nsdfjhsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949418,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\nsdfjhsdf\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949419,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\nsdfjhsdf\\n\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949419,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\nsdfjhsdf\\ns\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949419,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\nsdfjhsdf\\nsd\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949419,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\nsdfjhsdf\\nsdfh\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949419,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\nsdfjhsdf\\nsdfhj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949419,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\nsdfjhsdf\\nsdfhjj\",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949420,"sid":4,"ssid":40,"body":"[9411432970,\"dsfhgh\\nsdjfhj\\nshdjhj \\nsdhfj \\nsdhfj \\nsdfjhsdf\\nsdfhjj \",{\"f\":12,\"c\":\"#ff4b59\"}]"},{"ts":1511949425,"sid":4,"ssid":1,"body":"[[8,9425427430,[249.2857142857143,270.7142857142858],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949425,"sid":4,"ssid":1,"body":"[[8,9425889690,[280,279.28571428571433],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949431,"sid":4,"ssid":20,"body":"{\"ids\":[9209392590],\"offset\":{\"x\":41,\"y\":298}}"},{"ts":1511949434,"sid":4,"ssid":1,"body":"[[8,9434157420,[355.7142857142857,572.1428571428571],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949434,"sid":4,"ssid":1,"body":"[[8,9434641240,[366.42857142857144,581.4285714285714],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949437,"sid":7,"ssid":10,"body":"{\"key\":\"pageId\",\"value\":\"14e972b2d94e30f2571e2f033c0ca7f5.png\"}"},{"ts":1511949438,"sid":4,"ssid":1,"body":"[[8,9438320410,[207.85714285714286,242.85714285714286],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949439,"sid":4,"ssid":1,"body":"[[8,9439535220,[172.14285714285717,110.00000000000001],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949444,"sid":4,"ssid":10,"body":"[9117558090]"},{"ts":1511949445,"sid":4,"ssid":10,"body":"[9266453040]"},{"ts":1511949446,"sid":4,"ssid":10,"body":"[]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[1.4,61.599999999999994,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[18.2,67.19999999999999,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[26.599999999999998,70,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[32.199999999999996,74.19999999999999,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[43.4,78.39999999999999,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[50.4,85.39999999999999,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[54.599999999999994,88.19999999999999,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[64.39999999999999,93.8,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[67.19999999999999,96.6,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[78.39999999999999,103.6,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[84,107.8,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[96.6,112,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[100.8,113.39999999999999,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[103.6,113.39999999999999,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[106.39999999999999,113.39999999999999,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[107.8,112,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[107.8,110.6,\"\"]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[109.19999999999999,114.8,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[77.85714285714286,82.14285714285715]}]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[110.6,120.39999999999999,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[79.28571428571428,86.42857142857143]}]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[124.6,144.2,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[89.28571428571429,102.85714285714286]}]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[137.2,162.39999999999998,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[97.85714285714286,115.71428571428572]}]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[169.39999999999998,200.2,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[120.71428571428572,142.8571428571429]}]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[193.2,226.79999999999998,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[137.85714285714286,162.14285714285714]}]"},{"ts":1511949449,"sid":4,"ssid":22,"body":"[239.39999999999998,270.2,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[170.71428571428572,192.8571428571429]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[253.39999999999998,285.59999999999997,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[181.42857142857144,203.5714285714286]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[288.4,316.4,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[205.71428571428572,225.71428571428572]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[305.2,330.4,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[217.8571428571429,235.71428571428572]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[320.59999999999997,343,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[229.2857142857143,245.00000000000003]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[334.59999999999997,355.59999999999997,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[239.28571428571428,253.57142857142858]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[359.79999999999995,376.59999999999997,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[257.14285714285717,268.5714285714286]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[372.4,385,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[265.7142857142857,275.00000000000006]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[379.4,392,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[271.42857142857144,280]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[404.59999999999997,410.2,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[289.28571428571433,292.8571428571429]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[418.59999999999997,420,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[299.2857142857143,300.00000000000006]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[442.4,435.4,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[315.7142857142857,311.42857142857144]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[456.4,445.2,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[326.42857142857144,317.8571428571429]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[487.2,463.4,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[347.8571428571429,331.42857142857144]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[498.4,471.79999999999995,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[356.42857142857144,337.14285714285717]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[543.1999999999999,498.4,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[387.8571428571429,356.42857142857144]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[565.5999999999999,513.8,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[404.28571428571433,367.14285714285717]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[600.5999999999999,534.8,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[428.5714285714286,382.14285714285717]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[618.8,547.4,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[442.14285714285717,390.7142857142857]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[662.1999999999999,571.1999999999999,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[472.8571428571429,407.8571428571429]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[684.5999999999999,582.4,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[489.2857142857143,416.42857142857144]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[719.5999999999999,600.5999999999999,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[513.5714285714286,428.5714285714286]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[733.5999999999999,607.5999999999999,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[523.5714285714287,434.28571428571433]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[753.1999999999999,617.4,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[537.857142857143,440.7142857142857]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[757.4,620.1999999999999,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[541.4285714285716,442.8571428571429]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[771.4,627.1999999999999,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[550.7142857142857,447.8571428571429]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[778.4,631.4,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[555.7142857142857,450.7142857142858]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[789.5999999999999,635.5999999999999,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[564.2857142857143,454.28571428571433]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[795.1999999999999,638.4,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[567.8571428571429,456.42857142857144]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[802.1999999999999,642.5999999999999,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[572.8571428571429,458.5714285714286]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[803.5999999999999,642.5999999999999,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[573.5714285714286,459.28571428571433]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[805,644,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[575.0000000000001,460]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[806.4,644,{\"rectDown\":[77.14285714285715,78.57142857142858],\"rectEnd\":[575.7142857142857,460]}]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[805,642.5999999999999,\"\"]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[802.1999999999999,641.1999999999999,\"\"]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[791,635.5999999999999,\"\"]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[768.5999999999999,623,\"\"]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[712.5999999999999,597.8,\"\"]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[680.4,585.1999999999999,\"\"]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[571.1999999999999,543.1999999999999,\"\"]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[534.8,529.1999999999999,\"\"]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[424.2,481.59999999999997,\"\"]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[383.59999999999997,464.79999999999995,\"\"]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[281.4,417.2,\"\"]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[221.2,385,\"\"]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[125.99999999999999,330.4,\"\"]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[74.19999999999999,295.4,\"\"]"},{"ts":1511949450,"sid":4,"ssid":22,"body":"[8.399999999999999,244.99999999999997,\"\"]"},{"ts":1511949451,"sid":4,"ssid":22,"body":"[]"},{"ts":1511949452,"sid":4,"ssid":1,"body":"[[8,9452720390,[197.14285714285717,245.71428571428572],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949452,"sid":4,"ssid":22,"body":"[]"},{"ts":1511949454,"sid":4,"ssid":1,"body":"[[8,9454351660,[223.57142857142858,238.57142857142858],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949455,"sid":4,"ssid":1,"body":"[[8,9455634800,[205.71428571428572,269.2857142857143],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949456,"sid":4,"ssid":1,"body":"[[8,9456084630,[231.42857142857144,274.2857142857143],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949459,"sid":4,"ssid":1,"body":"[[8,9459830910,[423.5714285714286,296.42857142857144],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949462,"sid":4,"ssid":1,"body":"[[8,9462406780,[615,118.57142857142857],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949463,"sid":4,"ssid":1,"body":"[[8,9463782630,[349.2857142857143,234.2857142857143],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949466,"sid":4,"ssid":1,"body":"[[8,9466231410,[452.14285714285717,300.00000000000006],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949466,"sid":4,"ssid":1,"body":"[[8,9466659830,[360.7142857142858,239.28571428571428],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949468,"sid":4,"ssid":1,"body":"[[8,9468864370,[483.5714285714286,285],{\"c\":\"#ff4b59\",\"w\":5,\"f\":12}]]"},{"ts":1511949471,"sid":7,"ssid":90,"body":"[{\"key\":\"switch\",\"value\":-1}]"}]}`
)

func TestZlibCompress(t *testing.T) {
	t1 := text1

	out, err := ZlibCompress([]byte(t1))
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = ZlibUncompress(out)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestZlibSize(t *testing.T) {

	f, err := os.Create("/Users/dellinger/temp/origin.txt")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	f.Write([]byte(text3))
	f.Close()

	f, err = os.Create("/Users/dellinger/temp/comp.dat")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	buf, err := ZlibCompress([]byte(text3))
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	f.Write(buf)
	f.Close()
}
