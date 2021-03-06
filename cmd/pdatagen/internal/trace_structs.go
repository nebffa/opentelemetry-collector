// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

var traceFile = &File{
	Name: "trace",
	imports: []string{
		`otlpcommon "go.opentelemetry.io/collector/internal/data/opentelemetry-proto-gen/common/v1"`,
		`otlptrace "go.opentelemetry.io/collector/internal/data/opentelemetry-proto-gen/trace/v1"`,
	},
	testImports: []string{
		`"testing"`,
		``,
		`"github.com/stretchr/testify/assert"`,
		``,
		`otlptrace "go.opentelemetry.io/collector/internal/data/opentelemetry-proto-gen/trace/v1"`,
	},
	structs: []baseStruct{
		resourceSpansSlice,
		resourceSpans,
		instrumentationLibrarySpansSlice,
		instrumentationLibrarySpans,
		spanSlice,
		span,
		spanEventSlice,
		spanEvent,
		spanLinkSlice,
		spanLink,
		spanStatus,
	},
}

var resourceSpansSlice = &sliceStruct{
	structName: "ResourceSpansSlice",
	element:    resourceSpans,
}

var resourceSpans = &messageStruct{
	structName:     "ResourceSpans",
	description:    "// InstrumentationLibrarySpans is a collection of spans from a LibraryInstrumentation.",
	originFullName: "otlptrace.ResourceSpans",
	fields: []baseField{
		resourceField,
		&sliceField{
			fieldName:       "InstrumentationLibrarySpans",
			originFieldName: "InstrumentationLibrarySpans",
			returnSlice:     instrumentationLibrarySpansSlice,
		},
	},
}

var instrumentationLibrarySpansSlice = &sliceStruct{
	structName: "InstrumentationLibrarySpansSlice",
	element:    instrumentationLibrarySpans,
}

var instrumentationLibrarySpans = &messageStruct{
	structName:     "InstrumentationLibrarySpans",
	description:    "// InstrumentationLibrarySpans is a collection of spans from a LibraryInstrumentation.",
	originFullName: "otlptrace.InstrumentationLibrarySpans",
	fields: []baseField{
		instrumentationLibraryField,
		&sliceField{
			fieldName:       "Spans",
			originFieldName: "Spans",
			returnSlice:     spanSlice,
		},
	},
}

var spanSlice = &sliceStruct{
	structName: "SpanSlice",
	element:    span,
}

var span = &messageStruct{
	structName: "Span",
	description: "// Span represents a single operation within a trace.\n" +
		"// See Span definition in OTLP: https://github.com/open-telemetry/opentelemetry-proto/blob/master/opentelemetry/proto/trace/v1/trace.proto#L37",
	originFullName: "otlptrace.Span",
	fields: []baseField{
		traceIDField,
		spanIDField,
		traceStateField,
		parentSpanIDField,
		nameField,
		&primitiveTypedField{
			fieldName:       "Kind",
			originFieldName: "Kind",
			returnType:      "SpanKind",
			rawType:         "otlptrace.Span_SpanKind",
			defaultVal:      "SpanKindUNSPECIFIED",
			testVal:         "SpanKindSERVER",
		},
		startTimeField,
		endTimeField,
		attributes,
		droppedAttributesCount,
		&sliceField{
			fieldName:       "Events",
			originFieldName: "Events",
			returnSlice:     spanEventSlice,
		},
		&primitiveField{
			fieldName:       "DroppedEventsCount",
			originFieldName: "DroppedEventsCount",
			returnType:      "uint32",
			defaultVal:      "uint32(0)",
			testVal:         "uint32(17)",
		},
		&sliceField{
			fieldName:       "Links",
			originFieldName: "Links",
			returnSlice:     spanLinkSlice,
		},
		&primitiveField{
			fieldName:       "DroppedLinksCount",
			originFieldName: "DroppedLinksCount",
			returnType:      "uint32",
			defaultVal:      "uint32(0)",
			testVal:         "uint32(17)",
		},
		&messageField{
			fieldName:       "Status",
			originFieldName: "Status",
			returnMessage:   spanStatus,
		},
	},
}

var spanEventSlice = &sliceStruct{
	structName: "SpanEventSlice",
	element:    spanEvent,
}

var spanEvent = &messageStruct{
	structName: "SpanEvent",
	description: "// SpanEvent is a time-stamped annotation of the span, consisting of user-supplied\n" +
		"// text description and key-value pairs. See OTLP for event definition.",
	originFullName: "otlptrace.Span_Event",
	fields: []baseField{
		timeField,
		nameField,
		attributes,
		droppedAttributesCount,
	},
}

var spanLinkSlice = &sliceStruct{
	structName: "SpanLinkSlice",
	element:    spanLink,
}

var spanLink = &messageStruct{
	structName: "SpanLink",
	description: "// SpanLink is a pointer from the current span to another span in the same trace or in a\n" +
		"// different trace. See OTLP for link definition.",
	originFullName: "otlptrace.Span_Link",
	fields: []baseField{
		traceIDField,
		spanIDField,
		traceStateField,
		attributes,
		droppedAttributesCount,
	},
}

var spanStatus = &messageStruct{
	structName: "SpanStatus",
	description: "// SpanStatus is an optional final status for this span. Semantically when Status wasn't set\n" +
		"// it is means span ended without errors and assume Status.Ok (code = 0).",
	originFullName: "otlptrace.Status",
	fields: []baseField{
		&primitiveTypedField{
			fieldName:       "Code",
			originFieldName: "Code",
			returnType:      "StatusCode",
			rawType:         "otlptrace.Status_StatusCode",
			defaultVal:      "StatusCode(0)",
			testVal:         "StatusCode(1)",
		},
		&primitiveField{
			fieldName:       "Message",
			originFieldName: "Message",
			returnType:      "string",
			defaultVal:      `""`,
			testVal:         `"cancelled"`,
		},
	},
}

var traceIDField = &primitiveTypedField{
	fieldName:       "TraceID",
	originFieldName: "TraceId",
	returnType:      "TraceID",
	rawType:         "otlpcommon.TraceID",
	defaultVal:      "NewTraceID(nil)",
	testVal:         "NewTraceID([]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})",
}

var spanIDField = &primitiveTypedField{
	fieldName:       "SpanID",
	originFieldName: "SpanId",
	returnType:      "SpanID",
	rawType:         "otlpcommon.SpanID",
	defaultVal:      "NewSpanID(nil)",
	testVal:         "NewSpanID([]byte{1, 2, 3, 4, 5, 6, 7, 8})",
}

var parentSpanIDField = &primitiveTypedField{
	fieldName:       "ParentSpanID",
	originFieldName: "ParentSpanId",
	returnType:      "SpanID",
	rawType:         "otlpcommon.SpanID",
	defaultVal:      "NewSpanID(nil)",
	testVal:         "NewSpanID([]byte{8, 7, 6, 5, 4, 3, 2, 1})",
}

var traceStateField = &primitiveTypedField{
	fieldName:       "TraceState",
	originFieldName: "TraceState",
	returnType:      "TraceState",
	rawType:         "string",
	defaultVal:      `TraceState("")`,
	testVal:         `TraceState("congo=congos")`,
}

var droppedAttributesCount = &primitiveField{
	fieldName:       "DroppedAttributesCount",
	originFieldName: "DroppedAttributesCount",
	returnType:      "uint32",
	defaultVal:      "uint32(0)",
	testVal:         "uint32(17)",
}
