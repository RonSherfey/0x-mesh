package db

const v4OrdersSchema = `
CREATE TABLE IF NOT EXISTS ordersv4 (
	hash                     TEXT UNIQUE NOT NULL,
	chainID                  TEXT NOT NULL,
	verifyingContract        TEXT NOT NULL,
	makerToken               TEXT NOT NULL,
	takerToken               TEXT NOT NULL,
	makerAmount              TEXT NOT NULL,
	takerAmount              TEXT NOT NULL,
	takerTokenFeeAmount      TEXT NOT NULL,
	maker                    TEXT NOT NULL,
	taker                    TEXT NOT NULL,
	sender                   TEXT NOT NULL,
	feeRecipient             TEXT NOT NULL,
	pool                     TEXT NOT NULL,
	expiry                   TEXT NOT NULL,
	salt                     TEXT NOT NULL,
	signatureType            TEXT NOT NULL,
	signatureV               TEXT NOT NULL,
	signatureR               TEXT NOT NULL,
	signatureS               TEXT NOT NULL,
	lastUpdated              DATETIME NOT NULL,
	fillableTakerAssetAmount TEXT NOT NULL,
	isRemoved                BOOLEAN NOT NULL,
	isPinned                 BOOLEAN NOT NULL,
	isUnfillable             BOOLEAN NOT NULL,
	isExpired                BOOLEAN NOT NULL,
	lastValidatedBlockNumber TEXT NOT NULL,
	lastValidatedBlockHash   TEXT NOT NULL,
	keepCancelled            BOOLEAN NOT NULL,
	keepExpired              BOOLEAN NOT NULL,
	keepFullyFilled          BOOLEAN NOT NULL,
	keepUnfunded             BOOLEAN NOT NULL
);
`
const insertOrderQueryV4 = `INSERT INTO ordersv4 (
	hash,
	chainID,
	verifyingContract,
	makerToken,
	takerToken,
	makerAmount,
	takerAmount,
	takerTokenFeeAmount,
	maker,
	taker,
	sender,
	feeRecipient,
	pool,
	expiry,
	salt,
	signatureType,
	signatureV,
	signatureR,
	signatureS,
	lastUpdated,
	fillableTakerAssetAmount,
	isRemoved,
	isPinned,
	isUnfillable,
	isExpired,
	lastValidatedBlockNumber,
	lastValidatedBlockHash,
	keepCancelled,
	keepExpired,
	keepFullyFilled,
	keepUnfunded
) VALUES (
	:hash,
	:chainID,
	:verifyingContract,
	:makerToken,
	:takerToken,
	:makerAmount,
	:takerAmount,
	:takerTokenFeeAmount,
	:maker,
	:taker,
	:sender,
	:feeRecipient,
	:pool,
	:expiry,
	:salt,
	:signatureType,
	:signatureV,
	:signatureR,
	:signatureS,
	:lastUpdated,
	:fillableTakerAssetAmount,
	:isRemoved,
	:isPinned,
	:isUnfillable,
	:isExpired,
	:lastValidatedBlockNumber,
	:lastValidatedBlockHash,
	:keepCancelled,
	:keepExpired,
	:keepFullyFilled,
	:keepUnfunded
) ON CONFLICT DO NOTHING
`

const updateOrderQueryV4 = `UPDATE ordersv4 SET
	chainID = :chainID,
	verifyingContract = :verifyingContract,
	maker = :maker,
	makerToken = :makerToken,
	makerAmount = :makerAmount,
	taker = :taker,
	takerToken = :takerToken,
	takerAmount = :takerAmount,
	sender = :sender,
	feeRecipient = :feeRecipient,
	expiry = :expiry,
	salt = :salt,
	pool = :pool,
	signatureType = :signatureType,
	signatureV = :signatureV,
	signatureR = :signatureR,
	signatureS = :signatureS,
	lastUpdated = :lastUpdated,
	fillableTakerAssetAmount = :fillableTakerAssetAmount,
	isRemoved = :isRemoved,
	isPinned = :isPinned,
	isUnfillable = :isUnfillable,
	isExpired = :isExpired,
	lastValidatedBlockNumber = :lastValidatedBlockNumber,
	lastValidatedBlockHash = :lastValidatedBlockHash,
	keepCancelled = :keepCancelled,
	keepExpired = :keepExpired,
	keepFullyFilled = :keepFullyFilled,
	keepUnfunded = :keepUnfunded
WHERE ordersv4.hash = :hash
`