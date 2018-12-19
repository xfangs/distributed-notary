package proofs

import (
	"math/big"
	"reflect"

	"github.com/nkbai/goutils"
)

//证明拥有一个paillier私钥??
type NICorrectKeyProof struct {
	Sigma []*big.Int
}

const P = "1824183726245393467247644231302244136093537199057104213213550575243641782740360650490963459819244001947254814805714931325851267161710067435807383848463920901137710041673887113990643820798975386714299839914137592590654009952623014982962684955535111234311335565220917383311689115138496765310625882473439402233109450021984891450304679833752756159872219991089194187575068382762952239830901394850887215392132883640669486674102277756753260855640317510235617660944873631054630816035269100510337643250389997837634640249480037184290462924540150133678312185777228630834940021427688892876384196895677819059963882587092166301131529174343474451480089653483180602591751073139733370712300241581635049350925412683097729232092096276490229965785020041921736307394438075266234968515443716828633392848203945374591926800464450599823553052462708727219173990177119684565306222502415160037753326638045687574106534702341439991863742806351468290587722561435038912863815688133288619512790095919904026573249557024839383595481704184528960978957724597263323512030743875614290609368530643094080051166226135271385866188054556684837921935888945641944961066293159525602885452222458958772845494346799890196718717317906330936509091221354615991671869862034179206244894205681566781062633415772628848878715803040358836098609654889521393046492471227546079924219055408612815173193108753184477562256266860297096223934088509777393752624380757072082427603556077039945700711226680778392737267707541904355129695919972995501581794067880959822149963798096452613619855673307435602850208850402301583025111762622381953251883429317603005626232012725708694401272295509035367654620412640848204179955980722707996291909812529974361949926881288349518750747615837667549305083291804187179123453121466640918862622766511668478452223742058912575427337018022812631386313110243745000214354806312441270889672903307645611658893986526812130032112540367173736664288995222516688120866114984318582900331631896931709005163853429427759224323636152573453333607357348169167915027700846002932742550824939007414330697249569916339964247646402851281857942965519194576006169066153524163225631476643914033601957614124583206541834352791003930506139209204661104882701842617501635864883760885236797081996679496751254260706438583316885612406386543479255566185697792942478704336254208839180970748624881039948192415929866204318800220295457932550799088592217150597176505394120909914475575501881459804699385499576326684695531034075283165800622328491384987194944504461864105986907646706095956083156240472489616473946638879726524585936511018780747174387840018674670110430528051586069422163934697899931456041802624175449157279620104126331489491525955411465073551652840009163781923401029513048693746713122813578721687858104388238796796690"

var SALT_STRING = []byte{75, 90, 101, 110}

const M2 = 11
const DIGEST_SIZE = 256

func compute_digest(inputs ...[]byte) *big.Int {
	digest := utils.Sha256(inputs...)
	return new(big.Int).SetBytes(digest[:])
}

func CreateNICorrectKeyProof(key *PrivateKey) *NICorrectKeyProof {
	keyLength := key.n.BitLen()
	saltBn := new(big.Int).SetBytes(SALT_STRING)
	var rhos []*big.Int
	for i := 0; i < M2; i++ {
		seed := compute_digest(key.n.Bytes(), saltBn.Bytes(), big.NewInt(int64(i)).Bytes())
		mg := new(big.Int).Mod(mask_generation(keyLength, seed), key.n)
		rhos = append(rhos, mg)
	}
	var sigmas []*big.Int
	for i := 0; i < len(rhos); i++ {
		sigma := ExtractNroot(key, rhos[i])
		sigmas = append(sigmas, sigma)
	}
	return &NICorrectKeyProof{
		sigmas,
	}
}

//const
func (proof *NICorrectKeyProof) Verify(pubkey *PublicKey) bool {
	keyLength := pubkey.N.BitLen()
	saltBn := new(big.Int).SetBytes(SALT_STRING)
	var rho_vec []*big.Int
	for i := 0; i < M2; i++ {
		seed := compute_digest(pubkey.N.Bytes(), saltBn.Bytes(), big.NewInt(int64(i)).Bytes())
		mg := new(big.Int).Mod(mask_generation(keyLength, seed), pubkey.N)
		rho_vec = append(rho_vec, mg)
	}
	alpha_primorial, _ := new(big.Int).SetString(P, 10)
	gcd_test := new(big.Int).GCD(nil, nil, alpha_primorial, pubkey.N)
	var derived_rho_vec []*big.Int
	for i := 0; i < M2; i++ {
		tmp := new(big.Int).Exp(proof.Sigma[i], pubkey.N, pubkey.N)
		derived_rho_vec = append(derived_rho_vec, tmp)
	}
	if reflect.DeepEqual(rho_vec, derived_rho_vec) && gcd_test.Cmp(big.NewInt(1)) == 0 {
		return true
	}
	return false
}

// generate random element of size :
func mask_generation(outLength int, seed *big.Int) *big.Int {
	msklen := outLength/DIGEST_SIZE + 1
	var msklenHashs []*big.Int
	for j := 0; j < msklen; j++ {
		digest := compute_digest(seed.Bytes(), big.NewInt(int64(j)).Bytes())
		msklenHashs = append(msklenHashs, digest)
	}
	result := big.NewInt(0)
	for i := 0; i < len(msklenHashs); i++ {
		//log.Trace(fmt.Sprintf("result=%s", result.Text(16)))
		xtmp := msklenHashs[i]
		//log.Trace(fmt.Sprintf("xtmp=%s", xtmp.Text(16)))
		xtmp.Lsh(xtmp, uint(i*DIGEST_SIZE))
		result.Add(result, xtmp)
	}
	return result
}
