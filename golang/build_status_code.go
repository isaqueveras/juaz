//  Created by Isaque Veras on 03/15/24.
//  Copyright © 2024 Isaque Veras. All rights reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package golang

func _build_status_code(code string) string {
	switch code {
	case "100":
		return "http.StatusContinue"
	case "101":
		return "http.StatusSwitchingProtocols"
	case "102":
		return "http.StatusProcessing"
	case "103":
		return "http.StatusEarlyHints"
	case "200":
		return "http.StatusOK"
	case "201":
		return "http.StatusCreated"
	case "202":
		return "http.StatusAccepted"
	case "203":
		return "http.StatusNonAuthoritativeInfo"
	case "204":
		return "http.StatusNoContent"
	case "205":
		return "http.StatusResetContent"
	case "206":
		return "http.StatusPartialContent"
	case "207":
		return "http.StatusMultiStatus"
	case "208":
		return "http.StatusAlreadyReported"
	case "226":
		return "http.StatusIMUsed"
	case "300":
		return "http.StatusMultipleChoices"
	case "301":
		return "http.StatusMovedPermanently"
	case "302":
		return "http.StatusFound"
	case "303":
		return "http.StatusSeeOther"
	case "304":
		return "http.StatusNotModified"
	case "305":
		return "http.StatusUseProxy"
	case "307":
		return "http.StatusTemporaryRedirect"
	case "308":
		return "http.StatusPermanentRedirect"
	case "400":
		return "http.StatusBadRequest"
	case "401":
		return "http.StatusUnauthorized"
	case "402":
		return "http.StatusPaymentRequired"
	case "403":
		return "http.StatusForbidden"
	case "404":
		return "http.StatusNotFound"
	case "405":
		return "http.StatusMethodNotAllowed"
	case "406":
		return "http.StatusNotAcceptable"
	case "407":
		return "http.StatusProxyAuthRequired"
	case "408":
		return "http.StatusRequestTimeout"
	case "409":
		return "http.StatusConflict"
	case "410":
		return "http.StatusGone"
	case "411":
		return "http.StatusLengthRequired"
	case "412":
		return "http.StatusPreconditionFailed"
	case "413":
		return "http.StatusRequestEntityTooLarge"
	case "414":
		return "http.StatusRequestURITooLong"
	case "415":
		return "http.StatusUnsupportedMediaType"
	case "416":
		return "http.StatusRequestedRangeNotSatisfiable"
	case "417":
		return "http.StatusExpectationFailed"
	case "418":
		return "http.StatusTeapot"
	case "421":
		return "http.StatusMisdirectedRequest"
	case "422":
		return "http.StatusUnprocessableEntity"
	case "423":
		return "http.StatusLocked"
	case "424":
		return "http.StatusFailedDependency"
	case "425":
		return "http.StatusTooEarly"
	case "426":
		return "http.StatusUpgradeRequired"
	case "428":
		return "http.StatusPreconditionRequired"
	case "429":
		return "http.StatusTooManyRequests"
	case "431":
		return "http.StatusRequestHeaderFieldsTooLarge"
	case "451":
		return "http.StatusUnavailableForLegalReasons"
	case "500":
		return "http.StatusInternalServerError"
	case "501":
		return "http.StatusNotImplemented"
	case "502":
		return "http.StatusBadGateway"
	case "503":
		return "http.StatusServiceUnavailable"
	case "504":
		return "http.StatusGatewayTimeout"
	case "505":
		return "http.StatusHTTPVersionNotSupported"
	case "506":
		return "http.StatusVariantAlsoNegotiates"
	case "507":
		return "http.StatusInsufficientStorage"
	case "508":
		return "http.StatusLoopDetected"
	case "510":
		return "http.StatusNotExtended"
	case "511":
		return "http.StatusNetworkAuthenticationRequired"
	default:
		return "<nil>"
	}
}
