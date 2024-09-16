package layer_contract

import (
	"fmt"
	"strings"
)

// CircularReferenceException - <?php
//
// declare (strict_types=1);
// namespace Qossmic\Deptrac\Contract\Layer;
//
// use Qossmic\Deptrac\Contract\ExceptionInterface;
// use RuntimeException;
// use function implode;
// use function sprintf;
// /**
// * Exception when there are circular dependencies between layers.
// *
// * Thrown when you use the `layer_contract` collector and depend on a layer_contract that
// * in turn depends back on you. To be able to resolve layers, the dependencies
// * between them have to be a DAG(Direct Acyclic Graph), otherwise
// * the resolution is not possible.
// */
// final class CircularReferenceException extends RuntimeException implements ExceptionInterface
//
//	{
//	   /**
//	    * @param list<string> $others
//	    */
//	   public static function circularLayerDependency(string $layer_contract, array $others) : self
//	   {
//	       return new self(sprintf('Circular ruleset dependency_contract for layer_contract %s depending on: %s', $layer_contract, implode('->', $others)));
//	   }
//	}
type CircularReferenceException struct {
	Message string
}

func (c CircularReferenceException) Error() string {
	return c.Message
}

func NewCircularReferenceExceptionFromCircularLayerDependency(layer string, others []string) CircularReferenceException {
	return CircularReferenceException{Message: fmt.Sprintf("Circular ruleset dependency_contract for layer_contract %s depending on: %s", layer, strings.Join(others, "->"))}
}
